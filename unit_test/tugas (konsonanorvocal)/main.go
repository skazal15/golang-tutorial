package main

func KonsonOrVocal(s string) string {
	switch s {
	case "a", "i", "u", "e", "o":
		return "vocal"
	default:
		return "konsonan"
	}
}
