package main

const englishPrefix = "Greeting, "
const mandarinPrefix = "您好，"
const hawaiiPrefix = "Hola, "

func hello() string {
	return "Hello!"
}

func greeting(name string, locale string) string {

	if name == "" {
		name = "everyone"
	}

	prefix := englishPrefix

	switch locale {
	case "Mandarin":
		prefix = mandarinPrefix
	case "Hawaii":
		prefix = hawaiiPrefix
	}

	return prefix + name
}
