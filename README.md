# Wrapper do api dlisty w golangu

## Instalacja:

```
go get -u github.com/Minecrafterr/dlist-wrapper-golang
```

## Informacja
Aktualnie wrapper ma tylko kilka funkcji<br>
Będzie on stale rozwijany

## Przykładowe użycie
```go
package main
import "github.com/Minecrafterr/dlist-wrapper-golang"



func Main() {
// Updatowanie statystyk
  dlistwrapper.UpdateStats("token", liczba użytkowników, liczba serwerów)
// Funkcja wyświetlanie informacji o bocie. Aby wybrać dane które nas obchodzą wpiszcie na końcu funkcji np. .Data.Name po nazwę
var d dlistwrapper.GetBotInfo("ID bota z dblisty")
 fmt.Println("Nazwa bota: "+d.Data.Name+"\nOpis krótki: "+d.Info.ShortDescription)
 // Funkcja wyświetlania informacji o serwerze. Aby wybrać dane które nas obchodzą wpiszcie na końcu funkcji np. .Data.Name po nazwę
 var s dlistwrapper.GetServerInfo("id serwera z dblisty")
fmt.Println("Nazwa serwera: "+s.Data.Name+"\nOpis krótki: "+s.Info.ShortDescription)
}
```

UWAGA! Powyższy kod jest tylko przykładem!
## Dokumentacja wrappera
Dokumentacja wrappera jest w pliku docs.md
