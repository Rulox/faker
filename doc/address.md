# Address
Since version 1.0.0

This Generator will provide a full set of address. It can be used to get
single fields, or full addresses using the selected locale formatters.

Some fields don't make any sense in some Languages. For example, in Spanish, street type
comes before the street name (`Avenida <name>`) but in English is the other way around (`<name> Avenue`).
This means, some functions might not make sense in the selected locale. They'll simply return an
empty string to avoid any crash. 

```go
    AddressGenerator.City()         // Brooklyn
    AddressGenerator.Country()      // Spain 
    AddressGenerator.CountryCode()  // UK 
    AddressGenerator.StreetSuffix() // Calle 
    AddressGenerator.StreetPrefix() // Avenue 
    AddressGenerator.Street()       // Union 
    AddressGenerator.Province()     // Lyon
    AddressGenerator.State()        // New York
    AddressGenerator.Secondary()    // Stairs ##
    AddressGenerator.Number()       // 169 
    AddressGenerator.ZipCode()      // 10003 
    AddressGenerator.Full()         // 169 Bedford Ave, 11256 Brooklyn, New York
```

