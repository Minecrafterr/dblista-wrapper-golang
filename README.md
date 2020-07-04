# Wrapper do api dblisty w golangu
## Instalacja:

```
go get -u github.com/Minecrafterr/dblista-wrapper-golang
```

## Informacja
Aktualnie wrapper ma tylko 2 funkcje (aktualizowanie statystyk oraz informacje o bocie)
Będzie on stale rozwijany

## Przykładowe użycie
```go
package main
import "github.com/Minecrafterr/dblista-wrapper-golang"



func Main() {
// Updatowanie statystyk
  dblistawrapper.UpdateStats("token", "liczba użytkowników", "liczba serwerów")
// Funkcja wyświetlanie informacji o bocie. Aby wybrać dane które nas obchodzą wpiszcie na końcu funkcji np. .Data.Name po imię
var d dblistawrapper.GetBotInfo("ID bota z dblisty")
 fmt.Println("Nazwa bota: "+d.Data.Name+"\nOpis krótki: "+d.Info.ShortDescription)
}
```

UWAGA! Powyższy kod jest tylko przykładem!
