package dateparser_test

import (
	"fmt"
	"testing"

	dps "github.com/markusmobius/go-dateparser"
)

func BenchmarkParser_Search(b *testing.B) {
	p := dps.Parser{}
	for n := 0; n < b.N; n++ {
		for _, text := range benchmarkSearchTexts {
			_, _, err := p.Search(nil, text)
			if err != nil {
				fmt.Println(text)
				b.Fatal(err)
			}
		}
	}
}

var benchmarkSearchTexts = []string{
	`في 29 يوليو 1938 غزت القوات اليابانية الاتحاد السوفييتي ووقعت أولى المعارك والتي انتصر فيها السوفييت، وعلى الرغم من ذلك رفضت اليابان الاعتراف بذلك وقررت في 11 مايو 1939 تحريك الحدود المنغولية حتى نهر غول، حيث وقعت معركة خالخين غول والتي انتصر فيها الجيش الأحمر على جيش كوانتونغ`,
	`Пасля апублікавання Патсдамскай дэкларацыі 26 ліпеня 1945 года і адмовы Японіі капітуляваць на яе ўмовах ЗША скінулі атамныя бомбы.`,
	`На 16 юни 1944 г. започват въздушни бомбардировки срещу Япония, използувайки новозавладените острови като бази.`,
	`不過大多數人仍多把第二次世界大戰的爆發定為1939年9月1日德國入侵波蘭開始，這次入侵行動隨即導致英國與法國向德國宣戰。`,
	`V roce 1920 byla proto vytvořena Společnost národů, jež měla fungovat jako fórum, na němž měly národy mírovým způsobem urovnávat svoje spory.`,
	`Krigen i Europa begyndte den 1. september 1939, da Nazi-Tyskland invaderede Polen, og endte med Nazi-Tysklands betingelsesløse overgivelse den 8. maj 1945.`,
	` De meest dramatische uitbreiding van het conflict vond plaats op 22 juni 1941 met de Duitse aanval op de Sovjet-Unie.`,
	`I will meet you tomorrow at noon`,
	`in a minute`,
	// `last decade`,
	`July 13th.,	 July 14th`,
	`last updated Aug 06, 2018 05:05 PM CDT`,
	`25th march 2015 , i need this report today.`,
	`25th march 2015 , i need this report today.`,
	`Maraming namatay sa mga Hapon hanggang sila'y sumuko noong Agosto 15, 1945.`,
	`Iso-Britannia ja Ranska julistivat sodan Saksalle 3. syyskuuta 1939.`,
	`La 2e Guerre mondiale, ou Deuxième Guerre mondiale4, est un conflit armé à l'échelle planétaire qui dura du 1 septembre 1939 au 2 septembre 1945.`,
	`במרץ 1938 "אוחדה" אוסטריה עם גרמניה (אנשלוס). `,
	`जुलाई 1937 में, मार्को-पोलो ब्रिज हादसे का बहाना लेकर जापान ने चीन पर हमला कर दिया और चीनी साम्राज्य की राजधानी बीजिंग पर कब्जा कर लिया,`,
	`A háború Európában 1945. május 8-án Németország feltétel nélküli megadásával, míg Ázsiában szeptember 2-án, Japán kapitulációjával fejeződött be.`,
	`1937 წელს დაიწყო იაპონია-ჩინეთის მეორე ომი.`,
	`Die UdSSR blieb gemäß dem Neutralitätspakt vom 13. April 1941 gegenüber Japan vorerst neutral.`,
	`Kekaisaran Jepang menyerah pada tanggal 15 Agustus 1945, sehingga mengakhiri perang di Asia dan memperkuat kemenangan total Sekutu atas Poros.`,
	` Con questo il 2 ottobre 1935 prese il via la campagna d'Etiopia. Il 9 maggio 1936 venne proclamato l'Impero. `,
	`1939年9月1日、ドイツ軍がポーランドへ侵攻したことが第二次世界大戦の始まりとされている。`,
	`نگ جهانی دوم جنگ جدی بین سپتامبر 1939 و 2 سپتامبر 1945 بود.`,
	`II wojna światowa – największa wojna światowa w historii, trwająca od 1 września 1939 do 2 września 1945 (w Europie do 8 maja 1945)`,
	`Em outubro de 1936, Alemanha e Itália formaram o Eixo Roma-Berlim.`,
	`Pe 17 septembrie 1939, după semnarea unui acord de încetare a focului cu Japonia, sovieticii au invadat Polonia dinspre est.`,
	`Втора́я мирова́я война́ (1 сентября 1939 — 2 сентября 1945) — война двух мировых военно-политических коалиций, ставшая крупнейшим вооружённым конфликтом в истории человечества.`,
	`Desde finales de 1939 hasta inicios de 1941 Alemania conquistó o sometió gran parte de la Europa continental.`,
	`Efter kommunisternas seger 1922 drog de allierade och Japan bort sina trupper.`,
	`และเมื่อวันที่ 11 พฤษภาคม 1939 ญี่ปุ่นตัดสินใจขยายพรมแดนญี่ปุ่น-มองโกเลียขึ้นไปถึงแม่น้ำคัลคินกอลด้วยกำลัง`,
	`Almanya’nın Polonya’yı işgal ettiği 1 Eylül 1939 savaşın başladığı tarih olarak genel kabul görür.`,
	`Інші дати, що розглядаються деякими авторами як дати початку війни: початок японської інтервенції в Маньчжурію 13 вересня 1931, початок другої японсько-китайської війни 7 липня 1937 року та початок угорсько-української війни 14 березня 1939 року.`,
	`Ý theo gương Đức, đã tiến hành xâm lược Ethiopia năm 1935 và sát nhập Albania vào ngày 12 tháng 4 năm 1939.`,
	`January 3, 2017 - February 1st`,
	`2014 was good! October was excellent! Friday, 21 was especially good!`,
	`May 2020,	June 2020,	2023,	January UTC,	June 5 am utc,	June 23th 5 pm EST,	May 31, 8am UTC`,
	`19 марта 2001 был хороший день. 20 марта тоже был хороший день. 21 марта был отличный день.`,
	`19 марта 2001. Сегодня был хороший день. 2 дня назад был хороший день. Вчера тоже был хороший день.`,
	`19 марта 2001. Сегодня был хороший день. Два дня назад был хороший день. Хорошая была неделя. Думаю, через неделю будет еще лучше.`,
	`1962 augusztus 11 Föld körüli pályára bocsátották a szovjet Vosztok-3 űrhajót, mely páros űrrepülést hajtott végre a másnap föld körüli pályára bocsátott Vosztok-4-gyel.2 hónappal ezelőtt furcsa, nem forgó jellegű szédülést tapasztaltam.`,
	`1/1/1940. Vào tháng 8 năm 1940, với lực lượng lớn của Pháp tại Bắc Phi chính thức trung lập trong cuộc chiến, Ý mở một cuộc tấn công vào thuộc địa Somalia của Anh tại Đông Phi. Đến tháng 9 quân Ý vào đến Ai Cập (cũng đang dưới sự kiểm soát của Anh). `,
	`July 12th, 2014. July 13th, July 14th`,
	`2014. July 13th July 14th`,
	`July 13th 2014 July 14th 2014`,
	`July 13th 2014 July 14th`,
	`July 13th, 2014 July 14th, 2014`,
	`2014. July 12th, July 13th, July 14th`,
	`1938–1939 marscherade tyska soldater i Österrike samtidigt som österrikiska soldater marscherade i Berlin.`,
	`Verteidiger der Stadt kapitulierten am 2. Mai 1945. Am 8. Mai 1945 (VE-Day) trat bedingungslose Kapitulation der Wehrmacht in Kraft`,
	`19 марта 2001 был хороший день. 20 марта тоже был хороший день. 21 марта был отличный день.`,
	`Em outubro de 1936, Alemanha e Itália formaram o Eixo Roma-Berlim.`,
	`19 марта 2001, 20 марта, 21 марта был отличный день.`,
	``,
	// `Привет`,
	`DECEMBER 21 19.87 87`,
	`bonjour, pouvez vous me joindre svp par telephone 08 11 58 54 41`,
	`a Americ`,
	`9/3/2017  , `,
	`9/3/2017  ' `,
	`15 de outubro de 1936`,
	`19 марта 2001`,
}
