all:
	 go build main.go
clean:
	rm -rf main
grammar:
	 java -jar $(PWD)/bin/antlr-4.7-complete.jar -Dlanguage=Go -o parser Gherkin.g4 
