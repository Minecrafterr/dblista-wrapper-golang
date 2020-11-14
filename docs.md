# Dokumentacja wrappera
## WAŻNE! 27 Listopada zmiana nazwy wrappera (dblista-wrapper-golang -> dlist-wrapper-go)
## Dane
### Token api
W niektórych funkcjach do botów trzeba wziąść token api. Skąd można go zabrać?<br>
- W swoim bocie kliknij przycisk edytuj (jeśli jesteś właścicielem drugim itd. dopisz do linku bota /edit<br>
- Zjedź do samych przycisków Anuluj i Zapisz<br>
- Mamy kategorię statystyki bota (https://prnt.sc/tc8d70)<br>
- Klikamy wygeneruj token api i kopiujemy token do schowka (możesz już anulować edycję)<br>
### Token użytkownika
Aby uzyskać token użytkownika trzeba:<br>
- zalogować się na https://dlist.top<br>
- wejść w informacje o stronie -> pliki cookies<br>
- znaleźć klucz `sessionToken`<br>
- odkodować go z base64 (można użyć https://base64decode.org)<br>
## Boty
### Update statystyk
```go
dblistawrapper.UpdateStats("token", liczba użytkowników, liczba serwerów)
```
Token można wziąść z edycji bota
### Głosowanie na bota
```go
dblistawrapper.VoteBot("id bota", "token użytkownika")
```
### Dodawanie opinii
```go
dblistawrapper.RateBot("id bota", ocena 1-5, "komentarz", "token użytkownika")
```
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
### Ulepszanie bota (Wymagane premium na dbliście)
```go
dblistawrapper.BoostBot("id bota", "token użytkownika")
```
### Usuwanie ulepszenia bota (Wymagane premium na dbliście)
```go
dblistawrapper.RemoveBotBoost("id bota", "token użytkownika")
```
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
### Ulepszanie serwera (Wymagane premium na dbliście)
```go
dblistawrapper.BoostServer("id serwera", "token użytkownika")
```
### Usuwanie ulepszenia serwera (Wymagane premium na dbliście)
```go
dblistawrapper.RemoveServerBoost("id serwera", "token użytkownika")
```
### Głosowanie na serwer
```go
dblistawrapper.VoteServer("id serwera", "token użytkownika")
```
### Dodawanie opinii
```go
dblistawrapper.RateServer("id serwera", ocena 1-5, "komentarz", "token użytkownika")
```
## Użytkownik
### Informacje o użytkowniku
```go
dblistawrapper.getUserInfo("id użytkownika")
```
Lista dostępnych rzeczy (użytkownik):
.Data.Avatar - Avatar użytkownika
.Data.ID - ID użytkownika
.Data.LastRequestTime - Timestamp ostatniego requesta
.Data.RecentRequests - Ostatnie requesty
.Data.Username - Nazwa 
.Data.Ban - Czy zbanowany na dbliście
.Data.Money - Ilość monet
.Data.Perm - Poziom permisji
## Informacja
Wrapper jest ciągle rozwijany!
Autor: Minecrafter!#9481 - Nadmoderator DListy
