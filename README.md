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
import (
	"github.com/Minecrafterr/dblista-wrapper-golang"
)


func Main() {
// Updatowanie statystyk
  dblistawrapper.UpdateStats("token", "liczba użytkowników", "liczba serwerów")
// Wyświetlanie informacji o serwerze (wysyła do logów informacje o podanym bocie z dblisty)
 dblistawrapper.GetBotInfo("ID bota z dblisty")
}
```

UWAGA! Powyższy kod jest tylko przykładem!
