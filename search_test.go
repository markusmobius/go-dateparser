package dateparser_test

import (
	"fmt"
	"testing"
	"time"

	dps "github.com/markusmobius/go-dateparser"
	"github.com/stretchr/testify/assert"
)

func TestParser_SearchWithLanguage(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Language      string
		Text          string
		CurrentTime   time.Time
		ExtractedText []string
		ExtractedTime []time.Time
		ParserTypes   []dps.ParserType
	}

	utcNow := time.Now().UTC()
	utcDay, utcMonth := utcNow.Day(), int(utcNow.Month())

	tests := []testScenario{
		// ==========================
		// HAS CURRENT TIME SPECIFIED
		// ==========================

		{ // Arabic
			Language: "ar",
			Text: `في 29 يوليو 1938 غزت القوات اليابانية الاتحاد` +
				` السوفييتي ووقعت أولى المعارك والتي انتصر فيها السوفييت، وعلى الرغم من ذلك رفضت` +
				` اليابان الاعتراف بذلك وقررت في 11 مايو 1939 تحريك الحدود المنغولية حتى نهر غول،` +
				` حيث وقعت معركة خالخين غول والتي انتصر فيها الجيش الأحمر على جيش كوانتونغ`,
			CurrentTime: tt(2000, 1, 1),
			ExtractedText: []string{
				"في 29 يوليو 1938",
				"في 11 مايو 1939",
			},
			ExtractedTime: []time.Time{
				tt(1938, 7, 29),
				tt(1939, 5, 11),
			},
		}, { // Belarusian
			Language: "be",
			Text: `Пасля апублікавання Патсдамскай дэкларацыі 26 ліпеня 1945 года і адмовы Японіі капітуляваць ` +
				`на яе ўмовах ЗША скінулі атамныя бомбы.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"26 ліпеня 1945 года і"},
			ExtractedTime: []time.Time{tt(1945, 7, 26)},
		}, { // Bulgarian
			Language: "bg",
			Text: `На 16 юни 1944 г. започват въздушни ` +
				`бомбардировки срещу Япония, използувайки новозавладените острови като бази.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"На 16 юни 1944 г"},
			ExtractedTime: []time.Time{tt(1944, 6, 16)},
		}, { // Chinese
			Language:      "zh",
			Text:          `不過大多數人仍多把第二次世界大戰的爆發定為1939年9月1日德國入侵波蘭開始，這次入侵行動隨即導致英國與法國向德國宣戰。`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1939年9月1"},
			ExtractedTime: []time.Time{tt(1939, 9, 1)},
		}, { // Czech
			Language: "cs",
			Text: `V roce 1920 byla proto vytvořena Společnost národů, jež měla fungovat jako fórum, ` +
				`na němž měly národy mírovým způsobem urovnávat svoje spory.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1920"},
			ExtractedTime: []time.Time{tt(1920, 1, 1)},
		}, { // Danish
			Language: "da",
			Text: `Krigen i Europa begyndte den 1. september 1939, da Nazi-Tyskland invaderede Polen, ` +
				`og endte med Nazi-Tysklands betingelsesløse overgivelse den 8. maj 1945.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1. september 1939", "8. maj 1945"},
			ExtractedTime: []time.Time{tt(1939, 9, 1), tt(1945, 5, 8)},
		}, { // Dutch
			Language: "nl",
			Text: ` De meest dramatische uitbreiding van het conflict vond plaats op 22 juni 1941 met de ` +
				`Duitse aanval op de Sovjet-Unie.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"22 juni 1941"},
			ExtractedTime: []time.Time{tt(1941, 6, 22)},
		}, { // English
			Language:      "en",
			Text:          `I will meet you tomorrow at noon`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"tomorrow at noon"},
			ExtractedTime: []time.Time{tt(2000, 1, 2, 12)},
		}, {
			Language:      "en",
			Text:          `in a minute`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"in a minute"},
			ExtractedTime: []time.Time{tt(2000, 1, 1, 0, 1)},
		}, {
			Language:      "en",
			Text:          `last decade`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"last decade"},
			ExtractedTime: []time.Time{tt(1990, 1, 1)},
		}, {
			Language:      "en",
			Text:          "July 13th.\r\n July 14th",
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"July 13th", "July 14th"},
			ExtractedTime: []time.Time{tt(2000, 7, 13), tt(2000, 7, 14)},
		}, {
			Language:      "en",
			Text:          "last updated Aug 06, 2018 05:05 PM CDT",
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"Aug 06, 2018 05:05 PM CDT"},
			ExtractedTime: []time.Time{tt(2018, 8, 6, 17, 5, 18_000)}, // +18_000 second for CDT
		}, {
			Language:      "en",
			Text:          "25th march 2015 , i need this report today.",
			ExtractedText: []string{"25th march 2015"},
			ExtractedTime: []time.Time{tt(2015, 3, 25)},
			ParserTypes: []dps.ParserType{
				dps.Timestamp,
				dps.CustomFormat,
				dps.AbsoluteTime,
				dps.NoSpacesTime,
			},
		}, {
			Language:      "en",
			Text:          "25th march 2015 , i need this report today.",
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"25th march 2015", "today"},
			ExtractedTime: []time.Time{tt(2015, 3, 25), tt(2000, 1, 1)},
		}, {
			Language:      "en",
			Text:          "The employee has not submitted their documents till date",
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"till date"},
			ExtractedTime: []time.Time{tt(2000, 1, 1)},
		}, { // Filipino / Tagalog
			Language:      "tl",
			Text:          `Maraming namatay sa mga Hapon hanggang sila'y sumuko noong Agosto 15, 1945.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"noong Agosto 15, 1945"},
			ExtractedTime: []time.Time{tt(1945, 8, 15)},
		}, { // Finnish
			Language:      "fi",
			Text:          `Iso-Britannia ja Ranska julistivat sodan Saksalle 3. syyskuuta 1939.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"3. syyskuuta 1939"},
			ExtractedTime: []time.Time{tt(1939, 9, 3)},
		}, { // French
			Language: "fr",
			Text: `La 2e Guerre mondiale, ou Deuxième Guerre mondiale4, est un conflit armé à ` +
				`l'échelle planétaire qui dura du 1 septembre 1939 au 2 septembre 1945.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1 septembre 1939", "2 septembre 1945"},
			ExtractedTime: []time.Time{tt(1939, 9, 1), tt(1945, 9, 2)},
		}, { // Hebrew
			Language:      "he",
			Text:          `במרץ 1938 "אוחדה" אוסטריה עם גרמניה (אנשלוס). `,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"במרץ 1938"},
			ExtractedTime: []time.Time{tt(1938, 3, 1)},
		}, { // Hindi
			Language: "hi",
			Text: `जुलाई 1937 में, मार्को-पोलो ब्रिज हादसे का बहाना लेकर जापान ने चीन पर हमला कर दिया और चीनी साम्राज्य ` +
				`की राजधानी बीजिंग पर कब्जा कर लिया,`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"जुलाई 1937 में"},
			ExtractedTime: []time.Time{tt(1937, 7, 1)},
		}, { // Hungarian
			Language: "hu",
			Text: `A háború Európában 1945. május 8-án Németország feltétel nélküli megadásával, ` +
				`míg Ázsiában szeptember 2-án, Japán kapitulációjával fejeződött be.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1945. május 8-án", "szeptember 2-án"},
			ExtractedTime: []time.Time{tt(1945, 5, 8), tt(2000, 9, 2)},
		}, { // Georgian
			Language:      "ka",
			Text:          `1937 წელს დაიწყო იაპონია-ჩინეთის მეორე ომი.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1937"},
			ExtractedTime: []time.Time{tt(1937, 1, 1)},
		}, { // German
			Language: "de",
			Text: `Die UdSSR blieb gemäß dem Neutralitätspakt ` +
				`vom 13. April 1941 gegenüber Japan vorerst neutral.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"Die", "13. April 1941"},
			ExtractedTime: []time.Time{tt(1999, 12, 28), tt(1941, 4, 13)},
		}, { // Indonesian
			Language: "id",
			Text: `Kekaisaran Jepang menyerah pada tanggal 15 Agustus 1945, sehingga mengakhiri perang ` +
				`di Asia dan memperkuat kemenangan total Sekutu atas Poros.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"tanggal 15 Agustus 1945"},
			ExtractedTime: []time.Time{tt(1945, 8, 15)},
		}, { // Italian
			Language: "it",
			Text: ` Con questo il 2 ottobre 1935 prese il via la campagna ` +
				`d'Etiopia. Il 9 maggio 1936 venne proclamato l'Impero. `,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"2 ottobre 1935", "9 maggio 1936"},
			ExtractedTime: []time.Time{tt(1935, 10, 2), tt(1936, 5, 9)},
		}, { // Japanese
			Language:      "ja",
			Text:          `1939年9月1日、ドイツ軍がポーランドへ侵攻したことが第二次世界大戦の始まりとされている。`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1939年9月1"},
			ExtractedTime: []time.Time{tt(1939, 9, 1)},
		}, { // Persian
			Language:      "fa",
			Text:          `نگ جهانی دوم جنگ جدی بین سپتامبر 1939 و 2 سپتامبر 1945 بود.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"سپتامبر 1939", "2 سپتامبر 1945"},
			ExtractedTime: []time.Time{tt(1939, 9, 1), tt(1945, 9, 2)},
		}, { // Polish
			Language: "pl",
			Text: `II wojna światowa – największa wojna światowa w historii, ` +
				`trwająca od 1 września 1939 do 2 września 1945 (w Europie do 8 maja 1945)`,
			CurrentTime: tt(2000, 1, 1),
			ExtractedText: []string{
				"1 września 1939",
				"2 września 1945",
				"8 maja 1945",
			},
			ExtractedTime: []time.Time{
				tt(1939, 9, 1),
				tt(1945, 9, 2),
				tt(1945, 5, 8),
			},
		}, { // Portuguese
			Language:      "pt",
			Text:          `Em outubro de 1936, Alemanha e Itália formaram o Eixo Roma-Berlim.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"Em outubro de 1936"},
			ExtractedTime: []time.Time{tt(1936, 10, 1)},
		}, { // Romanian
			Language: "ro",
			Text: `Pe 17 septembrie 1939, după semnarea unui acord de încetare a focului cu Japonia, ` +
				`sovieticii au invadat Polonia dinspre est.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"17 septembrie 1939"},
			ExtractedTime: []time.Time{tt(1939, 9, 17)},
		}, { // Russian
			Language: "ru",
			Text: `Втора́я мирова́я война́ (1 сентября 1939 — 2 сентября 1945) — ` +
				`война двух мировых военно-политических коалиций, ставшая крупнейшим вооружённым ` +
				`конфликтом в истории человечества.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1 сентября 1939", "2 сентября 1945"},
			ExtractedTime: []time.Time{tt(1939, 9, 1), tt(1945, 9, 2)},
		}, { // Spanish
			Language: "es",
			Text: `Desde finales de 1939 hasta inicios de 1941 Alemania conquistó o sometió ` +
				`gran parte de la Europa continental.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"de 1939", "de 1941"},
			ExtractedTime: []time.Time{tt(1939, 1, 1), tt(1941, 1, 1)},
		}, {
			Language:      "es",
			Text:          `¡¡Ay!! En Madrid, a 17 de marzo de 1615. ¿Vos bueno?`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"a 17 de marzo de 1615"},
			ExtractedTime: []time.Time{tt(1615, 3, 17)},
		}, { // Swedish
			Language:      "sv",
			Text:          `Efter kommunisternas seger 1922 drog de allierade och Japan bort sina trupper.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1922"},
			ExtractedTime: []time.Time{tt(1922, 1, 1)},
		}, { // Thai
			Language: "th",
			Text: `และเมื่อวันที่ 11 พฤษภาคม 1939 ` +
				`ญี่ปุ่นตัดสินใจขยายพรมแดนญี่ปุ่น-มองโกเลียขึ้นไปถึงแม่น้ำคัลคินกอลด้วยกำลัง`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"11 พฤษภาคม 1939"},
			ExtractedTime: []time.Time{tt(1939, 5, 11)},
		}, { // Turkish
			Language: "tr",
			Text: `Almanya’nın Polonya’yı işgal ettiği 1 Eylül 1939 savaşın başladığı ` +
				`tarih olarak genel kabul görür.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"1 Eylül 1939"},
			ExtractedTime: []time.Time{tt(1939, 9, 1)},
		}, { // Ukrainian
			Language: "uk",
			Text: `Інші дати, що розглядаються деякими авторами як дати початку війни: початок японської ` +
				`інтервенції в Маньчжурію 13 вересня 1931, початок другої японсько-китайської війни 7 ` +
				`липня 1937 року та початок угорсько-української війни 14 березня 1939 року.`,
			CurrentTime: tt(2000, 1, 1),
			ExtractedText: []string{
				"13 вересня 1931",
				"7 липня 1937 року",
				"14 березня 1939 року",
			},
			ExtractedTime: []time.Time{
				tt(1931, 9, 13),
				tt(1937, 7, 7),
				tt(1939, 3, 14),
			},
		}, { // Vietnamese
			Language: "vi",
			Text: `Ý theo gương Đức, đã tiến hành xâm lược Ethiopia năm 1935 và sát ` +
				`nhập Albania vào ngày 12 tháng 4 năm 1939.`,
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"năm 1935", "ngày 12 tháng 4 năm 1939"},
			ExtractedTime: []time.Time{tt(1935, 1, 1), tt(1939, 4, 12)},
		},

		// =====================================================================
		// CURRENT TIME NOT SPECIFIED, SO CHECK THE ABILTY TO GENERATE BASE TIME
		// =====================================================================

		{ // English
			Language:      "en",
			Text:          `January 3, 2017 - February 1st`,
			ExtractedText: []string{"January 3, 2017", "February 1st"},
			ExtractedTime: []time.Time{tt(2017, 1, 3), tt(2017, 2, 1)},
		}, {
			Language: "en",
			Text: `2014 was good! October was excellent!` +
				` Friday, 21 was especially good!`,
			ExtractedText: []string{"2014", "October", "Friday, 21"},
			ExtractedTime: []time.Time{
				tt(2014, utcMonth, utcDay),
				tt(2014, 10, utcDay),
				tt(2014, 10, 21),
			},
		}, {
			Language: "en",
			Text: "May 2020\n" +
				"July 2020\n" +
				"2023\n" +
				"January UTC\n" +
				"June 5 am utc\n" +
				"June 23th 5 pm EST\n" +
				"May 31, 8am UTC",
			ExtractedText: []string{
				"May 2020", "July 2020",
				"2023", "January UTC",
				"June 5 am utc",
				"June 23th 5 pm EST",
				"May 31", "8am UTC",
			},
			ExtractedTime: []time.Time{
				tt(2020, 5, utcDay), tt(2020, 7, utcDay),
				tt(2023, 7, utcDay), tt(2023, 1, utcDay),
				tt(2023, 6, 5),
				tt(2023, 6, 23, 17, 0, 18_000), // +18_000 second for EST
				// Note for result below: in original the returned time is in UTC,
				// meanwhile here the returned time is in EST. I believe our behavior
				// is the correct one since the parser should use previous result as
				// base time.
				tt(2023, 5, 31, 0, 0, 18_000),
				tt(2023, 8, 31),
			},
		}, { // Russian
			Language: "ru",
			Text:     `19 марта 2001 был хороший день. 20 марта тоже был хороший день. 21 марта был отличный день.`,
			ExtractedText: []string{
				"19 марта 2001",
				"20 марта",
				"21 марта",
			},
			ExtractedTime: []time.Time{
				tt(2001, 3, 19),
				tt(2001, 3, 20),
				tt(2001, 3, 21),
			},
		}, {
			Language: "ru",
			Text:     `Андрей Дмитриевич Сахаров скончался 14 декабря в 1989 году от внезапной остановки сердца.`,
			ExtractedText: []string{
				"14 декабря в 1989 году",
			},
			ExtractedTime: []time.Time{
				tt(1989, 12, 14, 0, 0),
			},
		}, { // Russian with relative dates
			Language: "ru",
			Text: `19 марта 2001. Сегодня был хороший день. 2 дня назад был хороший день. ` +
				`Вчера тоже был хороший день.`,
			ExtractedText: []string{
				"19 марта 2001",
				"Сегодня",
				"2 дня назад",
				"Вчера",
			},
			ExtractedTime: []time.Time{
				tt(2001, 3, 19),
				tt(2001, 3, 19),
				tt(2001, 3, 17),
				tt(2001, 3, 18),
			},
		}, {
			Language: "ru",
			Text: `19 марта 2001. Сегодня был хороший день. Два дня назад был хороший день. Хорошая была неделя. ` +
				`Думаю, через неделю будет еще лучше.`,
			ExtractedText: []string{
				"19 марта 2001",
				"Сегодня",
				"Два дня назад",
				"через неделю",
			},
			ExtractedTime: []time.Time{
				tt(2001, 3, 19),
				tt(2001, 3, 19),
				tt(2001, 3, 17),
				tt(2001, 3, 26),
			},
		}, { // Hungarian
			Language: "hu",
			Text: `1962 augusztus 11 Föld körüli pályára bocsátották a szovjet Vosztok-3 űrhajót, ` +
				`mely páros űrrepülést hajtott végre a másnap föld körüli pályára bocsátott Vosztok-4-gyel.` +
				`2 hónappal ezelőtt furcsa, nem forgó jellegű szédülést tapasztaltam.`,
			ExtractedText: []string{"1962 augusztus 11", "2 hónappal ezelőtt"},
			ExtractedTime: []time.Time{tt(1962, 8, 11), tt(1962, 6, 11)},
		}, { // Vietnamese
			Language: "vi",
			Text: `1/1/1940. Vào tháng 8 năm 1940, với lực lượng lớn của Pháp tại Bắc Phi chính thức trung lập ` +
				`trong cuộc chiến, Ý mở một cuộc tấn công vào thuộc địa Somalia của Anh tại Đông Phi. ` +
				`Đến tháng 9 quân Ý vào đến Ai Cập (cũng đang dưới sự kiểm soát của Anh). `,
			ExtractedText: []string{
				"1/1/1940",
				"tháng 8 năm 1940",
				"tháng 9",
			},
			ExtractedTime: []time.Time{
				tt(1940, 1, 1),
				tt(1940, 8, 1),
				tt(1940, 9, 1),
			},
		},

		// ====================================
		// CHECK FUNCTION "SPLIT IF NOT PARSED"
		// ====================================

		{ // English
			Language: "en",
			Text:     "July 12th, 2014. July 13th, July 14th",
			ExtractedText: []string{
				"July 12th, 2014",
				"July 13th",
				"July 14th",
			},
			ExtractedTime: []time.Time{
				tt(2014, 7, 12),
				tt(2014, 7, 13),
				tt(2014, 7, 14),
			},
		}, {
			Language: "en",
			Text:     "2014. July 13th July 14th",
			ExtractedText: []string{
				"2014",
				"July 13th",
				"July 14th",
			},
			ExtractedTime: []time.Time{
				tt(2014, utcMonth, utcDay),
				tt(2014, 7, 13),
				tt(2014, 7, 14),
			},
		}, {
			Language:      "en",
			Text:          "July 13th 2014 July 14th 2014",
			ExtractedText: []string{"July 13th 2014", "July 14th 2014"},
			ExtractedTime: []time.Time{tt(2014, 7, 13), tt(2014, 7, 14)},
		}, {
			Language:      "en",
			Text:          "July 13th 2014 July 14th",
			ExtractedText: []string{"July 13th 2014", "July 14th"},
			ExtractedTime: []time.Time{tt(2014, 7, 13), tt(2014, 7, 14)},
		}, {
			Language:      "en",
			Text:          "July 13th, 2014 July 14th, 2014",
			ExtractedText: []string{"July 13th, 2014", "July 14th, 2014"},
			ExtractedTime: []time.Time{tt(2014, 7, 13), tt(2014, 7, 14)},
		}, {
			Language: "en",
			Text:     "2014. July 12th, July 13th, July 14th",
			ExtractedText: []string{
				"2014",
				"July 12th",
				"July 13th",
				"July 14th",
			},
			ExtractedTime: []time.Time{
				tt(2014, utcMonth, utcDay),
				tt(2014, 7, 12),
				tt(2014, 7, 13),
				tt(2014, 7, 14),
			},
		}, { // Swedish
			Language: "sv",
			Text: "1938–1939 marscherade tyska soldater i Österrike samtidigt som " +
				"österrikiska soldater marscherade i Berlin.",
			ExtractedText: []string{"1938", "1939"},
			ExtractedTime: []time.Time{tt(1938, utcMonth, utcDay), tt(1939, utcMonth, utcDay)},
		}, { // German
			Language: "de",
			Text: "Verteidiger der Stadt kapitulierten am 2. Mai 1945. Am 8. Mai 1945 (VE-Day) trat " +
				"bedingungslose Kapitulation der Wehrmacht in Kraft",
			ExtractedText: []string{"am 2. Mai 1945", "Am 8. Mai 1945"},
			ExtractedTime: []time.Time{tt(1945, 5, 2), tt(1945, 5, 8)},
		},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare message
		message := fmt.Sprintf("%s \"%s\"", test.Language, test.Text)

		// Search for dates
		p := dps.Parser{ParserTypes: test.ParserTypes}
		cfg := dps.Configuration{CurrentTime: test.CurrentTime}
		result, err := p.SearchWithLanguage(&cfg, test.Language, test.Text)

		// Assert result
		passed := assert.NoError(t, err, message)
		if passed {
			passed = assert.Len(t, result, len(test.ExtractedTime), message)
		}
		if passed {
			for i, res := range result {
				textOK := assert.Equal(t, test.ExtractedText[i], res.Text, message)
				timeOK := assert.Zero(t, test.ExtractedTime[i].Sub(res.Date.Time).Seconds(), message)

				passed = textOK && timeOK
				if !passed {
					fmt.Printf("\t\t\tWANT \"%v\" GOT \"%v\" FROM \"%s\"\n",
						test.ExtractedTime[i], res.Date.Time, res.Text)
					break
				}
			}
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Search(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Languages        []string
		Text             string
		CurrentTime      time.Time
		StrictParsing    bool
		ExtractedText    []string
		ExtractedTime    []time.Time
		ExpectedLanguage string
	}

	tests := []testScenario{
		{ // Normal search
			Languages: []string{"en", "ru"},
			Text:      "19 марта 2001 был хороший день. 20 марта тоже был хороший день. 21 марта был отличный день.",
			ExtractedText: []string{
				"19 марта 2001",
				"20 марта",
				"21 марта",
			},
			ExtractedTime: []time.Time{
				tt(2001, 3, 19),
				tt(2001, 3, 20),
				tt(2001, 3, 21),
			},
		}, {
			Text:          "Em outubro de 1936, Alemanha e Itália formaram o Eixo Roma-Berlim.",
			CurrentTime:   tt(2000, 1, 1),
			ExtractedText: []string{"Em outubro de 1936"},
			ExtractedTime: []time.Time{tt(1936, 10, 1)},
		}, {
			Languages: []string{"en", "ru"},
			Text:      "19 марта 2001, 20 марта, 21 марта был отличный день.",
			ExtractedText: []string{
				"19 марта 2001",
				"20 марта",
				"21 марта",
			},
			ExtractedTime: []time.Time{
				tt(2001, 3, 19),
				tt(2001, 3, 20),
				tt(2001, 3, 21),
			},
		}, { // Dates not found
			Text: "",
		}, { // Language not detected
			Languages: []string{"en"},
			Text:      "Привет",
		}, { // Zero division error
			Text:          "DECEMBER 21 19.87 87",
			ExtractedText: []string{"DECEMBER 21 19"},
			ExtractedTime: []time.Time{tt(2019, 12, 21)},
		}, {
			Text:          "bonjour, pouvez vous me joindre svp par telephone 08 11 58 54 41",
			StrictParsing: true,
		}, {
			Text: "a Americ",
		}, { // Date with comma and apostrophe
			Languages:     []string{"en"},
			Text:          "9/3/2017  , ",
			ExtractedText: []string{"9/3/2017"},
			ExtractedTime: []time.Time{tt(2017, 9, 3)},
		}, {
			Languages:     []string{"en"},
			Text:          "9/3/2017  ' ",
			ExtractedText: []string{"9/3/2017"},
			ExtractedTime: []time.Time{tt(2017, 9, 3)},
		}, { // With expected language
			Text:             "15 de outubro de 1936",
			ExtractedText:    []string{"15 de outubro de 1936"},
			ExtractedTime:    []time.Time{tt(1936, 10, 15)},
			ExpectedLanguage: "pt",
		}, { // Invalid language
			Languages: []string{"unknown language code"},
			Text:      "19 марта 2001",
		},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare message
		message := fmt.Sprintf("%v \"%s\"", test.Languages, test.Text)

		// Search for dates
		lang, result, _ := dps.Search(&dps.Configuration{
			Languages:     test.Languages,
			CurrentTime:   test.CurrentTime,
			StrictParsing: test.StrictParsing,
		}, test.Text)

		// Assert result
		passed := assert.Len(t, result, len(test.ExtractedTime), message)
		if passed && test.ExpectedLanguage != "" {
			passed = assert.Equal(t, test.ExpectedLanguage, lang, message)
		}
		if passed {
			for i, res := range result {
				textOK := assert.Equal(t, test.ExtractedText[i], res.Text, message)
				timeOK := assert.Zero(t, test.ExtractedTime[i].Sub(res.Date.Time).Seconds(), message)

				passed = textOK && timeOK
				if !passed {
					fmt.Printf("\t\t\tWANT \"%v\" GOT \"%v\" FROM \"%s\"\n",
						test.ExtractedTime[i], res.Date.Time, res.Text)
					break
				}
			}
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_SearchWithConfig(t *testing.T) {
	// Initial search for Arabic arabicText with zero config
	arabicText := `في 29 يوليو 1938 غزت القوات اليابانية الاتحاد`
	lang, result, _ := dps.Search(nil, arabicText)
	assert.Equal(t, "ar", lang)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, tt(1938, 7, 29), result[0].Date.Time)

	// Now we explicitly specify its locale to English, so the search failed.
	cfg := &dps.Configuration{Locales: []string{"en-UK"}}
	_, result, _ = dps.Search(cfg, arabicText)
	assert.Empty(t, result)

	// Here we explicitly specify its language to English so the search still failed.
	cfg = &dps.Configuration{Languages: []string{"en"}}
	_, result, _ = dps.Search(cfg, arabicText)
	assert.Empty(t, result)

	// When both locale and language specified, language is ignored
	cfg = &dps.Configuration{Locales: []string{"ar"}, Languages: []string{"en"}}
	_, result, _ = dps.Search(cfg, arabicText)
	assert.Equal(t, "ar", lang)
	assert.Equal(t, tt(1938, 7, 29), result[0].Date.Time)

	cfg = &dps.Configuration{Locales: []string{"en"}, Languages: []string{"ar"}}
	_, result, _ = dps.Search(cfg, arabicText)
	assert.Empty(t, result)

	// Now we use custom detect language function that always return "en".
	// Since our text is Arabic, now the search will be failed.
	p := dps.Parser{
		DetectLanguagesFunction: func(string) []string {
			return []string{"en"}
		},
	}

	_, result, _ = p.Search(cfg, arabicText)
	assert.Empty(t, result)

	// When both language and language function detector specified,
	// the external function is ignored.
	cfg = &dps.Configuration{Languages: []string{"ar"}}
	_, result, _ = dps.Search(cfg, arabicText)
	assert.Equal(t, "ar", lang)
	assert.Equal(t, tt(1938, 7, 29), result[0].Date.Time)
}
