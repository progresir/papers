package papers

import (
	"fmt"
	"testing"
)

func TestBedOtherCountSymbolSNILS(t *testing.T) {
	snils := []string{
		`1585863071q42`,
		`150713`,
		`1585,в86`,
		`1585"63074313`,
	}
	for i := range snils {
		err := ValidateSnils(snils[i])
		if err != ErrSnils {
			t.Errorf("Snils verification error: %s", snils[i])
		}
	}
}

func TestBedOtherSymbolSNILS(t *testing.T) {
	snils := []string{
		`1585863071q`,
		`1585w630713`,
		`1585,в86013`,
		`1585"64b307`,
	}
	for i := range snils {
		err := ValidateSnils(snils[i])
		if err != ErrSnilsCotainInvalidCharacters {
			t.Errorf("Snils verification error: %s", snils[i])
		}
	}
}

func TestBedSNILS(t *testing.T) {
	snils := []string{
		`10389946`,
		`29`,
		`15858630`,
		`998443543543603`,
		`2106356`,
		``,
	}
	fmt.Println()
	for i := range snils {
		outputfunc := ValidateSnils(snils[i])
		if outputfunc != ErrSnils {
			t.Errorf("Snils verification error: %s", snils[i])
		}
	}
}

func TestWellSNILS(t *testing.T) {
	snils := []string{
		`10389946082`,
		`29706232998`,
		`15858630713`,
		`99848333603`,
		`21068746356`,
		`56643715619`,
		`47526472917`,
		`49471818840`,
		`57535921021`,
		`21783971904`,
		`875728553`,
		`52837223291`,
		`72374448711`,
		`173230212`,
		`51959698561`,
		`98932310352`,
		`64576679660`,
		`09998182564`,
		`31832480866`,
		`78106074698`,
		`84340732897`,
		`11816809866`,
		`22957103273`,
		`37955562445`,
		`33343222533`,
		`42752143365`,
		`71889342551`,
		`15562352767`,
		`26404930058`,
		`35253856387`,
		`96714379962`,
		`32150001587`,
		`20580959574`,
		`40165121217`,
		`84597171659`,
		`79930463862`,
		`26626948012`,
		`47631295000`,
		`94087602527`,
		`76259117625`,
		`96756307971`,
		`79811794871`,
		`26319671084`,
		`69711104400`,
		`39549250630`,
		`63603786401`,
		`68846324557`,
		`29330251457`,
		`613147181`,
		`43575798334`,
		`73436787027`,
		`22768247698`,
		`16166755890`,
		`78639264266`,
		`14186875997`,
		`69582246251`,
		`60046409142`,
		`57010636758`,
		`99203649034`,
		`81698246551`,
		`79384384470`,
		`33924807495`,
		`68151843008`,
		`46986871280`,
		`79147154430`,
		`79737328170`,
		`44022990355`,
		`23059226547`,
	}

	for i := range snils {
		outputfunc := ValidateSnils(snils[i])
		if outputfunc != nil {
			t.Errorf("Snils verification error: %s", snils[i])
		}
	}
}

func TestINN(t *testing.T) {
	inn := []string{
		`369488401131`,
		`909826069600`,
		`015786111254`,
		`918584075780`,
		`109483472008`,
		`546438788272`,
		`186687008760`,
		`779097202431`,
		`105980679340`,
		`679206457302`,
		`187146492718`,
		`515737181703`,
		`519145171550`,
		`346359870585`,
		`057659704109`,
		`578012677230`,
		`795915479433`,
		`529933226489`,
		`567785238546`,
		`025358033398`,
		`186814102912`,
		`639235728431`,
		`405751148988`,
		`659564040690`,
		`257425289854`,
		`276289878672`,
		`916238968030`,
		`288240096804`,
		`428824529665`,
		`187763382900`,
		`318076938830`,
		`639123799829`,
		`465379973543`,
		`137219593701`,
		`575248993802`,
		`579044074903`,
		`525568873171`,
		`458446634570`,
		`146697265104`,
		`248710787250`,
		`136042161242`,
		`458592930792`,
		`469904880756`,
		`327273557073`,
		`329682016939`,
		`017953483378`,
		`809965669834`,
		`769552483907`,
		`145880721393`,
		`129382754707`,
	}
	for i := range inn {
		err := ValidateInn(inn[i])
		if err != nil {
			t.Errorf("INN verification error: %s", inn[i])
		}
	}
}

func TestBadINN(t *testing.T) {
	inn := []string{
		`369401131`,
		`9098260+9600`,
		`01578611Ъ254`,
		`918584аку5780`,
		`109483472432008`,
		`54643ак8788272`,
		`1866870куца08760`,
		``,
	}
	for i := range inn {
		err := ValidateInn(inn[i])
		if err == nil {
			t.Errorf("INN verification error: %s", inn[i])
		}
	}
}
