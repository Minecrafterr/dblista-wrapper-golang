# Dokumentacja wrappera
## Boty
### Update statystyk
```go
dblistawrapper.UpdateStats("token", "liczba użytkowników", "liczba serwerów")
```
Token można wziąść z edycji bota
### Głosowanie na bota
```go
dblistawrapper.VoteBot("id bota", "token użytkownika")
```
Aby uzyskać token użytkownika trzeba:<br>
- zalogować się na https://dblista.pl<br>
- wejść w informacje o stronie -> pliki cookies<br>
- znaleźć klucz `sessionToken`<br>
- odkodować go z base64 (można użyć https://base64decode.org)<br>
### Dodawanie opinii
```go
dblistawrapper.RateBot("id bota", ocena, "komentarz", "token użytkownika")
```
Aby uzyskać token użytkownika trzeba:<br>
- zalogować się na https://dblista.pl<br>
- wejść w informacje o stronie -> pliki cookies<br>
- znaleźć klucz `sessionToken`<br>
- odkodować go z base64 (można użyć https://base64decode.org)<br>
### Informacje o bocie
```go
// Wszystkie dane w jsonie
dblistawrapper.GetBotInfo("id bota")
// Konkretne dane w tym wypadku nazwa bota
dblistawrapper.GetBotInfo("id bota").Data.Name
```
Lista dostępnych rzeczy (boty):<br>
.Data.Name - Nazwa <br>
.Data.Avatar - Link do avataru<br>
.Data.ID - ID bota<br>
.Data.Owner - ID właściciela bota<br>
.Data.Votes - Liczba głosów<br>
.Data.InvitePermissions - Permisje w liczbie przy invite linku<br>
.Data.Votelog - Array głosów (timestamp oraz id użytkownika)<br>
.Data.Info.FullDescription - Długi opis<br>
.Data.Info.ShortDescription - Krótki opis<br>
.Data.Info.Library - Biblioteka bota<br>
.Data.Info.Prefix - Prefix bota<br>
.Data.Info.Tags - Tagi bota (w arrayu)<br>
.Data.Links.Discord - Link do discorda<br>
.Data.Links.Git - Repozytorium bota<br>
.Data.Links.Www - Strona bota<br>
.Data.Ratings.Avarage - Średnia ocen opinii<br>
.Data.Ratings.Order - Łączna ocena opinii<br>
.Data.Ratings.Ratings - Array opinii (id autora, treść, ocena)<br>
.Data.Stats.Members - Ilość użytkowników<br>
.Data.Stats.MonthlyInvites - Miesięcznych dodać<br>
.Data.Stats.Servers - Ilość serwerów<br>
.Data.Stats.Status - Status bota (online, offline, etc)<br>
.Data.Status.Verification - Status weryfikacji<br>
.Data.Uptime.Max - Maks punktów<br>
.Data.Uptime.Online - Ilość punktów kiedy bot był online<br>
## Serwery
### Informacje o serwerze
```go
// Wszystkie dane w jsonie
dblistawrapper.GetServerInfo("id")
// Konkretne dane w tym wypadku nazwa serwera
dblistawrapper.GetServerInfo("id").Data.Name
```
Lista dostępnych rzeczy (serwery):<br>
.Data.Name - Nazwa <br>
.Data.Avatar - Link do avataru<br>
.Data.ID - ID serwera<br>
.Data.Owner - ID właściciela serwera<br>
.Data.Votes - Liczba głosów<br>
.Data.Votelog - Array głosów (timestamp oraz id użytkownika)<br>
.Data.Info.FullDescription - Długi opis<br>
.Data.Info.ShortDescription - Krótki opis<br>
.Data.Info.Tags - Tagi bota (w arrayu)<br>
.Data.Links.Discord - Link do discorda<br>
.Data.Links.Www - Strona bota<br>
.Data.Ratings.Avarage - Średnia ocen opinii<br>
.Data.Ratings.Order - Łączna ocena opinii<br>
.Data.Ratings.Ratings - Array opinii (id autora, treść, ocena)<br>
.Data.Stats.MonthlyInvites - Miesięcznych dodać<br>
.Data.Uptime.Max - Liczba wszystkich użytkowników<br>
.Data.Uptime.Online - Liczba użytkowników online<br>
## Informacja
Wrapper jest ciągle rozwijany!
Autor: Minecrafter!#9481 - Moderator dblisty
