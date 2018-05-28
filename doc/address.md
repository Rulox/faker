# Address
Since version 1.0.0

This Generator will provide a full set of address. It can be used to get
single fields, or full addresses using the selected locale formatters.

Some fields don't make any sense in some Languages. For example, in Spanish, street type
comes before the street name (Avenida <name>) but in English is the other way around (<name> Avenue).
This means, some functions might not make sense in the selected locale. They'll simply return an
empty string to avoid any crash. 

```go
    AddressGenerator.Address.City()         // Brooklyn
    AddressGenerator.Address.Country()      // Spain 
    AddressGenerator.Address.CountryCode()  // UK 
    AddressGenerator.Address.StreetSuffix() // Calle 
    AddressGenerator.Address.StreetPrefix() // Avenue 
    AddressGenerator.Address.Street()       // Union 
    AddressGenerator.Address.Province()     // Lyon
    AddressGenerator.Address.State()        // New York
    AddressGenerator.Address.Secondary()    // Stairs ##
    AddressGenerator.Address.Number()       // 169 
    AddressGenerator.Address.ZipCode()      // 10003 
    AddressGenerator.Address.Full()         // 169 Bedford Ave, 11256 Brooklyn, New York
```

