package language

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestDetectLanguage(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Language string
		Text     string
	}

	tests := []testScenario{
		// Arabic
		{"ar", `في 29 يوليو 1938 غزت القوات اليابانية الاتحاد` +
			` السوفييتي ووقعت أولى المعارك والتي انتصر فيها السوفييت، وعلى الرغم من ذلك رفضت` +
			` اليابان الاعتراف بذلك وقررت في 11 مايو 1939 تحريك الحدود المنغولية حتى نهر غول،`},
		// Belarusian
		{"be", `Пасля апублікавання Патсдамскай дэкларацыі 26 ліпеня 1945 года і адмовы Японіі капітуляваць ` +
			`на яе ўмовах ЗША скінулі атамныя бомбы.`},
		// Bulgarian
		{"bg", `На 16 юни 1944 г. започват въздушни ` +
			`бомбардировки срещу Япония, използувайки новозавладените острови като бази.`},
		// Chinese
		{"zh", `不過大多數人仍多把第二次世界大戰的爆發定為1939年9月1日德國入侵波蘭開始，2015年04月08日10点05。`},
		// Czech
		{"cs", `V rok 1920 byla proto vytvořena Společnost národů, jež měla fungovat jako fórum, ` +
			`na němž měly národy mírovým způsobem urovnávat svoje spory.`},
		// Danish
		{"da", `Krigen i Europa begyndte den 1. september 1939, da Nazi-Tyskland invaderede Polen, ` +
			`og endte med Nazi-Tysklands betingelsesløse overgivelse den 8. marts 1945.`},
		// Dutch
		{"nl", ` De meest dramatische uitbreiding van het conflict vond plaats op Maandag 22 juni 1941  met de ` +
			`Duitse aanval op de Sovjet-Unie.`},
		// English
		{"en", `I will meet you tomorrow at noon`},
		// Filipino / Tagalog
		{"tl", `Maraming namatay sa mga Hapon hanggang sila'y sumuko noong Agosto 15, 1945.`},
		// Finnish
		{"fi", `Iso-Britannia ja Ranska julistivat sodan Saksalle 3. syyskuuta 1939.`},
		// French
		{"fr", `La Seconde Guerre mondiale, ou Deuxième Guerre mondiale4, est un conflit armé à ` +
			`l'échelle planétaire qui dura du 1 septembre 1939 au 2 septembre 1945.`},
		// Hebrew
		{"he", `במרץ 1938 "אוחדה" אוסטריה עם גרמניה (אנשלוס). `},
		// Hindi
		{"hi", `जुलाई 1937 में, मार्को-पोलो ब्रिज हादसे का बहाना लेकर जापान ने चीन पर हमला कर दिया और चीनी साम्राज्य ` +
			`की राजधानी बीजिंग पर कब्जा कर लिया,`},
		// Hungarian
		{"hu", `A háború Európában 1945. május 8-án Németország feltétel nélküli megadásával, ` +
			`míg Ázsiában szeptember 2-án, Japán kapitulációjával fejeződött be.`},
		// Georgian
		{"ka", `1937 წელს დაიწყო იაპონია-ჩინეთის მეორე ომი.`},
		// German
		{"de", `Die UdSSR blieb dem Neutralitätspakt ` +
			`vom 13. April 1941 gegenüber Japan vorerst neutral.`},
		// Indonesian
		{"id", `Kekaisaran Jepang menyerah pada tanggal 15 Agustus 1945, sehingga mengakhiri perang ` +
			`di Asia dan memperkuat kemenangan total Sekutu atas Poros.`},
		// Italian
		{"it", ` Con questo il 2 ottobre 1935 prese il via la campagna ` +
			`d'Etiopia. Il 9 maggio 1936 venne proclamato l'Impero. `},
		// Japanese
		{"ja", `1933年（昭和8年）12月23日午前6時39分、宮城（現：皇居）内の産殿にて誕生。`},
		// Persian
		{"fa", `نگ جهانی دوم جنگ جدی بین سپتامبر 1939 و 2 سپتامبر 1945 بود.`},
		// Polish
		{"pl", `II wojna światowa – największa wojna światowa w historii, ` +
			`trwająca od 1 września 1939 do 2 września 1945 (w Europie do 8 maja 1945)`},
		// Portuguese
		{"pt", `Em outubro de 1936, Alemanha e Itália formaram o Eixo Roma-Berlim.`},
		// Romanian
		{"ro", `Pe 17 septembrie 1939, după semnarea unui acord de încetare a focului cu Japonia, ` +
			`sovieticii au invadat Polonia dinspre est.`},
		// Russian
		{"ru", `Втора́я мирова́я война́ (1 сентября 1939 — 2 сентября 1945) — ` +
			`война двух мировых военно-политических коалиций, ставшая крупнейшим вооружённым ` +
			`конфликтом в истории человечества.`},
		// Spanish
		{"es", `11 junio 2010`},
		// Swedish
		{"sv", ` den 15 augusti 1945 då Kejsardömet`},
		// Thai
		{"th", `และเมื่อวันที่ 11 พฤษภาคม 1939 ` +
			`ญี่ปุ่นตัดสินใจขยายพรมแดนญี่ปุ่น-มองโกเลียขึ้นไปถึงแม่น้ำคัลคินกอลด้วยกำลัง`},
		// Turkish
		{"tr", `Almanya’nın Polonya’yı işgal ettiği 1 Eylül 1939 savaşın başladığı ` +
			`tarih olarak genel kabul görür.`},
		// Ukrainian
		{"uk", `Інші дати, що розглядаються деякими авторами як дати початку війни: початок японської ` +
			`інтервенції в Маньчжурію 13 вересня 1931, початок другої японсько-китайської війни 7 ` +
			`липня 1937 року та початок угорсько-української війни 14 березня 1939 року.`},
		// Vietnamese
		{"vi", `Ý theo gương Đức, đã tiến hành xâm lược Ethiopia năm 1935 và sát ` +
			`nhập Albania vào ngày 12 tháng 4 năm 1939.`},
		// Only digits
		{"en", `2007`},
	}

	// Prepare config
	cfg := setting.Configuration{SkipTokens: []string{"t"}}
	languages, _ := GetLanguages(nil, nil, false)
	uniqueCharsets := GetUniqueCharsets(languages)

	// Start test
	var nFailed int
	for _, test := range tests {
		bestLanguage, _ := DetectLanguage(&cfg, test.Text, languages, uniqueCharsets, nil)
		if passed := assert.Equal(t, test.Language, bestLanguage); !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
