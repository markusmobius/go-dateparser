"""Export native calendar data and parity cases from pinned Python dependencies."""

import argparse
import hashlib
import importlib
import importlib.metadata
import json
from datetime import date, datetime
from functools import cache
from itertools import product
from pathlib import Path

from convertdate import gregorian, persian
from dateparser.calendars.hijri_parser import hijri_parser
from dateparser.calendars.jalali_parser import jalali_parser
from dateparser.conf import settings
from dateparser.search import search_dates
from dateparser.search.search import DateSearchWithDetection, _add_time_span_results
from hijridate import Gregorian, Hijri, ummalqura


ROOT = Path(__file__).resolve().parents[2]
VERSIONS = {
    "dateparser": "1.4.3",
    "convertdate": "2.4.1",
    "hijridate": "2.6.0",
    "pymeeus": "0.5.12",
}
EPOCH = gregorian.to_jd(1970, 1, 1)


def calendar_data():
    for package, version in VERSIONS.items():
        installed = importlib.metadata.version(package)
        if installed != version:
            raise RuntimeError(f"Expected {package} {version}, got {installed}")

    sources = {}
    for name in [
        "dateparser.calendars",
        "dateparser.calendars.jalali_parser",
        "dateparser.calendars.hijri_parser",
        "convertdate.persian",
        "convertdate.gregorian",
        "pymeeus.Epoch",
        "pymeeus.Sun",
        "pymeeus.Earth",
        "hijridate.convert",
        "hijridate.ummalqura",
    ]:
        module = importlib.import_module(name)
        sources[name] = hashlib.sha256(Path(module.__file__).read_bytes()).hexdigest()

    equinox = persian.equinox_jd
    persian.equinox_jd = cache(equinox)
    try:
        first_year = persian.from_gregorian(1, 1, 1)[0]
        year_starts = []
        year = first_year
        while True:
            try:
                start = persian.to_jd(year, 1, 1) - EPOCH
            except ValueError:
                break
            if not start.is_integer():
                raise RuntimeError(f"Nonintegral Jalali day: {year}: {start}")
            year_starts.append(int(start))
            year += 1

        max_day = int(gregorian.to_jd(3000, 1, 1) - EPOCH)
        while True:
            parts = gregorian.from_jd(EPOCH + max_day + 1)
            try:
                persian.from_gregorian(*parts)
            except ValueError:
                break
            max_day += 1

        for index, start in enumerate(year_starts):
            year = first_year + index
            if year >= 1:
                parts = persian.to_gregorian(year, 1, 1)
                assert int(gregorian.to_jd(*parts) - EPOCH) == start
                assert persian.from_gregorian(*parts) == (year, 1, 1)
            if index + 1 < len(year_starts):
                length = year_starts[index + 1] - start
                assert length in (365, 366)
                assert persian.month_length(year, 12) == length - 336
    finally:
        persian.equinox_jd = equinox

    month_starts = [start + 2400000 - Gregorian(1970, 1, 1).to_julian() for start in ummalqura.MONTH_STARTS]
    hijri_first = ummalqura.HIJRI_RANGE[0][0]
    for index, start in enumerate(month_starts[:-1]):
        year, month = hijri_first + index // 12, index % 12 + 1
        reference = Hijri(year, month, 1)
        assert reference.month_length() == month_starts[index + 1] - start
        assert (reference.to_gregorian() - date(1970, 1, 1)).days == start

    return {
        "reference": {
            "module": "https://github.com/scrapinghub/dateparser",
            "commit": "9ce60b1958f1b285886bcfbb743f6419feacfc92",
            "versions": VERSIONS,
            "source_sha256": sources,
        },
        "jalali": {
            "first_year": first_year,
            "year_starts": year_starts,
            "gregorian_min": (date(1, 1, 1) - date(1970, 1, 1)).days,
            "gregorian_max": max_day,
        },
        "hijri": {
            "first_year": hijri_first,
            "month_starts": month_starts,
            "gregorian_min": month_starts[0],
            "gregorian_max": month_starts[-1] - 1,
        },
    }


def python_settings(configuration):
    values = {
        "RELATIVE_BASE": datetime.fromisoformat(configuration["current_time"])
    }
    for source, target in {
        "date_order": "DATE_ORDER",
        "strict_parsing": "STRICT_PARSING",
        "required_parts": "REQUIRE_PARTS",
        "return_time_as_period": "RETURN_TIME_AS_PERIOD",
        "ignore_surrounding_text": "IGNORE_SURROUNDING_TEXT",
        "skip_tokens": "SKIP_TOKENS",
        "default_languages": "DEFAULT_LANGUAGES",
        "use_given_order": "USE_GIVEN_LANGUAGE_ORDER",
        "return_time_span": "RETURN_TIME_SPAN",
        "default_start_of_week": "DEFAULT_START_OF_WEEK",
        "default_days_in_month": "DEFAULT_DAYS_IN_MONTH",
    }.items():
        if source in configuration:
            values[target] = configuration[source]
    for source, target, choices in [
        ("preferred_date_source", "PREFER_DATES_FROM", ["current_period", "past", "future"]),
        ("preferred_day_of_month", "PREFER_DAY_OF_MONTH", ["current", "first", "last"]),
        ("preferred_month_of_year", "PREFER_MONTH_OF_YEAR", ["current", "first", "last"]),
    ]:
        if source in configuration:
            values[target] = choices[configuration[source]]
    return settings.replace(mod_settings=values, **values)


def feature_data(reference, calendar_sha256):
    cases = []
    base = {"current_time": "2000-01-01T00:00:00Z", "preserve_end_of_month": True}

    def add(stage, text, configuration):
        case = {"id": f"python-{stage}-{len(cases) + 1:05}", "stage": stage,
                "input": text, "configuration": configuration, "matches": []}
        parser = jalali_parser if stage == "jalali" else hijri_parser
        try:
            value, period = parser.parse(text, python_settings(configuration))
            elapsed = value.replace(tzinfo=None) - datetime(1970, 1, 1)
            case["matches"].append({"text": text, "date": {
                "parsed": True, "unix_seconds": elapsed.days * 86400 + elapsed.seconds,
                "nanosecond": value.microsecond * 1000, "offset": 0,
                "period": period.title(), "timezone": "UTC",
            }})
        except Exception as error:
            case["error"] = f"{type(error).__name__}: {error}"
        cases.append(case)

    equinox = persian.equinox_jd
    persian.equinox_jd = cache(equinox)
    try:
        for year, month, day in product(
            [0, 1, 9, 38, 199, 426, 686, 756, 818, 1111, 1181, 1210,
             1300, 1355, 1356, 1387, 1399, 1400, 1402, 1403, 1444, 1600,
             1634, 1635, 2000, 2097, 2192, 2262, 2300, 2377, 2378, 2379, 3177, 3178],
            [1, 2, 6, 7, 11, 12], [1, 29, 30, 31],
        ):
            add("jalali", f"{day}/{month}/{year:04}", {**base, "date_order": "DMY"})
        for year, month, day in product(range(1343, 1501), range(1, 13), [1, 29, 30]):
            add("hijri", f"{day}/{month}/{year}", {**base, "date_order": "DMY"})
        inputs = ["", " \t ", "today", "year", "00", "10", "01/02", "1444", "01/1444",
                  "Farvardin 1444", "farvardin 1444", "1 Farvardin 1403", "1/1/1444",
                  "1444/1/2", "1/1/89", "1/1/90", "30/12/1444", "30/02/1444",
                  "31/1/1403", "1/1/1444 16:50 pm", "1/1/1444 09:10:11.123456789",
                  "1/1/1444 03:05 EST", "1/1/1444 UTC", "1/1/1444 GMT+05:30",
                  "1/1/1444 +05:30", "23:30", "24:00", "1/1/1444 unknown"]
        configurations = [base]
        configurations += [{**base, "date_order": order} for order in ["DMY", "DYM", "MDY", "MYD", "YMD", "YDM"]]
        configurations += [{**base, **extra} for extra in [
            {"strict_parsing": True}, {"required_parts": ["day"]},
            {"required_parts": ["day", "month", "year"]}, {"return_time_as_period": True},
            {"ignore_surrounding_text": True},
        ]]
        for stage, configuration, text in product(["jalali", "hijri"], configurations, inputs):
            add(stage, text, configuration)
        for source, month, day in product(range(3), repeat=3):
            configuration = {**base, "preferred_date_source": source,
                             "preferred_month_of_year": month, "preferred_day_of_month": day}
            for stage, text in product(["jalali", "hijri"], ["year", "01/02", "1444", "30/12/1444", "23:30"]):
                add(stage, text, configuration)
        for reference_time in ["0001-01-01T12:00:00Z", "1924-07-31T12:00:00Z", "1924-08-01T12:00:00Z",
                               "2024-03-20T12:00:00Z", "2077-11-16T12:00:00Z", "2077-11-17T12:00:00Z",
                               "2999-12-31T12:00:00Z", "3000-12-31T12:00:00Z"]:
            for stage, text in product(["jalali", "hijri"], ["1/1/1444", "01/02", "23:30"]):
                add(stage, text, {"current_time": reference_time})
        for year, month in product([1, 1342, 1501, 1502, 9999], [1, 12]):
            add("hijri", f"1/{month}/{year:04}", {**base, "date_order": "DMY"})
    finally:
        persian.equinox_jd = equinox

    search_reference = DateSearchWithDetection()

    def add_search(text, configuration, language=None, detector_languages=None):
        stage = "search_with_language" if language else "search"
        case = {"id": f"python-{stage}-{len(cases) + 1:05}", "stage": stage,
                "input": text, "configuration": configuration, "matches": []}
        if language:
            case["language"] = language
        detector = None
        if detector_languages is not None:
            case["detector_languages"] = detector_languages
            case["detection_inputs"] = []

            def detector(text, confidence_threshold=None):
                case["detection_inputs"].append(text)
                return detector_languages

        try:
            selected_settings = python_settings(configuration)
            strategy = configuration.get("search_strategy", "split")
            if language:
                if strategy == "ngram":
                    found = search_reference.ngram_search.search_parse([language], text, selected_settings)
                    found = _add_time_span_results(found, text, selected_settings)
                else:
                    found = search_reference.search.search_parse(language, text, selected_settings)
                found = [(text, value, language) for text, value in found]
            else:
                found = search_dates(text, languages=configuration.get("languages"),
                                     settings=selected_settings, strategy=strategy,
                                     add_detected_language=True, detect_languages_function=detector) or []
            for original, value, detected in found:
                offset = int(value.utcoffset().total_seconds()) if value.utcoffset() is not None else 0
                elapsed = value.replace(tzinfo=None) - datetime(1970, 1, 1)
                case["detected"] = detected
                case["matches"].append({"text": original, "date": {
                    "parsed": True, "unix_seconds": elapsed.days * 86400 + elapsed.seconds - offset,
                    "nanosecond": value.microsecond * 1000, "offset": offset,
                }})
        except Exception as error:
            case["error"] = f"{type(error).__name__}: {error}"
            if isinstance(error, IndexError):
                case["known_difference"] = "python-search-exception"
            elif isinstance(error, OverflowError):
                case["known_difference"] = "python-relative-range"
        cases.append(case)

    search_inputs = [
        ("en", "The first artificial Earth satellite was launched on 4 October 1957."),
        ("en", "Invoice issued February 5th, 2020; due March 1, 2020."),
        ("en", "from May 5, 2019 to June 7, 2019"),
        ("en", "on twenty first of March 2014"),
        ("en", "March 2014 CET"), ("en", "now PST"), ("en", "xyzz"),
        ("en", "Report #123 at 9:30 on 2024-02-29"),
        ("fr", "La r\u00e9union a eu lieu le 12 mars 2024."),
        ("de", "Am 12. M\u00e4rz 2024. Danach am 13. M\u00e4rz 2024."),
        ("ru", "\u0441 12 \u043c\u0430\u0440\u0442\u0430 2024 \u043f\u043e 14 \u043c\u0430\u0440\u0442\u0430 2024"),
        ("vi", "ng\u00e0y 12 th\u00e1ng 3 n\u0103m 2024"),
        ("hu", "2024. m\u00e1rcius 12."),
        ("zh", "2024-02-29 12:34:56 UTC"), ("ja", "2024-02-29 12:34:56 UTC"),
        ("yue", "2024-02-29 12:34:56 UTC"),
    ]
    for strategy in ["split", "ngram"]:
        configuration = {**base, "search_strategy": strategy}
        for language, text in search_inputs:
            add_search(text, configuration, language)
            add_search(text, {**configuration, "languages": [language]})
        for text in [search_inputs[0][1], search_inputs[8][1], "Meeting on 2024-02-29"]:
            add_search(text, configuration)
        for languages in [["fr", "en"], ["en", "fr"]]:
            add_search("Invoice issued February 5th, 2020", {**configuration, "languages": languages, "use_given_order": True})
        for detector in [["fr"], [], ["de", "en"]]:
            add_search("La r\u00e9union du 12 mars 2024", {**configuration, "default_languages": ["fr"]}, detector_languages=detector)
        add_search("March 2014 CET", {**configuration, "languages": ["en", "de"], "skip_tokens": ["xyzz"]})
        add_search("xyzz", {**configuration, "languages": ["en", "de"], "skip_tokens": ["xyzz"]})
        for week_start, month_days, expression in product(["monday", "sunday"], [28, 30], [
            "for the past month", "last week", "past 2 days", "past 2 weeks", "past 1 months",
            "next month", "next week", "next 2 days", "next 2 weeks", "next 1 months",
            "past 0 days", "next 2 weeks and last week", "past 99999999999999999999 weeks",
        ]):
            add_search("Messages received " + expression, {
                **configuration, "current_time": "2025-01-31T12:34:56.123456Z", "languages": ["en"],
                "return_time_span": True, "default_start_of_week": week_start,
                "default_days_in_month": month_days,
            })

    reference = {**reference, "versions": {**reference["versions"]}, "source_sha256": {**reference["source_sha256"]}}
    for package, version in {"python-dateutil": "2.9.0.post0", "pytz": "2026.3.post1", "regex": "2026.9.3",
                             "six": "1.17.0", "tzdata": "2026.3", "tzlocal": "5.4.4"}.items():
        if importlib.metadata.version(package) != version:
            raise RuntimeError(f"Expected {package} {version}")
        reference["versions"][package] = version
    for name in ["dateparser.conf", "dateparser.parser", "dateparser.date", "dateparser.languages.locale",
                 "dateparser.search.search", "dateparser.search.ngram_search", "dateparser.search.text_detection",
                 "dateparser.utils.time_spans", "dateparser.freshness_date_parser"]:
        module = importlib.import_module(name)
        reference["source_sha256"][name] = hashlib.sha256(Path(module.__file__).read_bytes()).hexdigest()
    return {"reference": reference, "calendar_data_sha256": calendar_sha256, "cases": cases}


def write_document(path, document, check):
    contents = (json.dumps(document, indent=2, ensure_ascii=True) + "\n").encode()
    if check:
        if path.read_bytes() != contents:
            raise RuntimeError(f"Reference data did not reproduce: {path}")
    else:
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_bytes(contents)
    return hashlib.sha256(contents).hexdigest()


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, default=ROOT / "internal/parser/calendars/data.json")
    parser.add_argument("--features-output", type=Path)
    parser.add_argument("--check", action="store_true")
    args = parser.parse_args()
    document = calendar_data()
    digest = write_document(args.output, document, args.check)
    print(f"Verified {len(document['jalali']['year_starts'])} Jalali year starts and "
          f"{len(document['hijri']['month_starts'])} Umm al-Qura month starts; "
          f"SHA256 {digest}")
    if args.features_output:
        features = feature_data(document["reference"], digest)
        digest = write_document(args.features_output, features, args.check)
        print(f"Exported {len(features['cases'])} Python feature cases; SHA256 {digest}")


if __name__ == "__main__":
    main()