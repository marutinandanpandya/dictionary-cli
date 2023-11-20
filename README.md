# dictionary-cli: A Golang Package to get definitions.

dictionary-cli is a golang package cli to get definitions of words from merriam-webster dictionary.

## Examples

```
go build -o ./dictionary .


export DICTIONARY_API_KEY="508c71a6-04b8-4cd9-abd8-XXXXXXXXX"
./dictionary apple
ˈa-pəl (noun): the fleshy, usually rounded red, yellow, or green edible pome fruit of a usually cultivated tree (genus Malus) of the rose family; also : an apple tree
./dictionary exercise
ˈek-sər-ˌsīz (noun): the act of bringing into play or realizing in action : use
```

When the input word is not found in dictionary
```
./dictionary excersice
Error: Requested word does not exists in dictionary.
Possible sugestions ["excursive","excessive","excursion","excelsior","excursively","exceptive","excursions","excessing","excursuses","excessively","excursus","exceedance","excrescence","exceedence","excursiveness","excellence","excrescences","excelsiors","excrescent","excerpting"]
```

