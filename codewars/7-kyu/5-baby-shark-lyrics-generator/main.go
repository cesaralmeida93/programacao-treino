package kata

const S = ", doo doo doo doo doo doo\n"

func BabySharkLyrics() (r string) {

	a := func(p string) { r += p + S + p + S + p + S + p + "!\n" }
	b := func(p string) { a(p + " shark") }

	b("Baby")
	b("Mommy")
	b("Daddy")
	b("Grandma")
	b("Grandpa")
	a("Let's go hunt")

	return r + "Run away,…\n"
}

// package kata
// import s "strings"
// func BabySharkLyrics() (r string) {
// 	z := s.Repeat
// 	for _, i := range []string{"Baby shark", "Mommy shark", "Daddy shark", "Grandma shark", "Grandpa shark", "Let's go hunt"} {
// 		r += z(i+","+z(" doo", 6)+"\n", 3) + i + "!\n"
// 	}
// 	return r + "Run away,…\n"
// }

// package kata
// import"strings"
// func BabySharkLyrics()string{
// t:="B,B,B,B!M,M,M,M!D,D,D,D!Gm,Gm,Gm,Gm!Gp,Gp,Gp,Gp!L,L,L,L!R"
// for _,R:=range strings.Split(`!
// $Ds$ma$pa$a s$, d d d d d d
// $Run away,…
// $doo$Baby s$Mommy s$shark$Grand$Daddy $Let's go hunt`,"$"){t=strings.ReplaceAll(t,R[:1],R)}
// return t}
