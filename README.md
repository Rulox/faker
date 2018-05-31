
<p align="center"><img src="doc/faker_logo.png" width="250"></p>

# Faker (WIP, do not use this library in production)
![quality](https://sonarcloud.io/api/project_badges/measure?project=faker_key&metric=alert_status) 
[![Build Status](https://travis-ci.org/Rulox/faker.svg?branch=master)](https://travis-ci.org/Rulox/faker)
[![GoDoc](https://godoc.org/github.com/Rulox/faker?status.svg)](https://godoc.org/github.com/Rulox/faker)

Faker is a Golang library that generates all type of fake data. Including localized data.

This library is inspired by same libs in other languages like Perl's, ruby's and PHP's Faker.

Faker has been built and tested with Go >= 1.10

## Content
- [Installing](#installing)
- [Usage](#usage)
    - [Misc](doc/misc.md)
    - [Address](doc/address.md)
    - [Person](doc/person.md)
- [Locales](#locales)
    - [Set Locale](#set-your-locale)
- [Examples](examples)
- [FAQ](#faq)
- [License](#license)

### Installing
Just use go get

`go get -u github.com/rulox/faker`

Or add the library to your project and use `dep ensure`

### Usage
The main struct `generator/Faker` provides all the usability for you in order to create the fake data.

`Faker` would have different number of generators (for Misc data, Addresses, Companies, Phones, etc). 

```go
	f := faker.NewFaker("en_US") // Use en_US locale
    f.Misc.RandomInt()  // 54
    f.Address.Street()  // Bedford 
``` 
### Locales
The locales are organized in YAML files called `faker.yml` inside each language folder.
The default locale is `en_US`. A lot of help to create new locales for languages is needed
and the format is so easy that anyone (even if you're not a developer) can supply data
following the format.

#### Set your locale
Setting a locale is really easy. The locale string is passed in the `NewFaker` function.
This function will (among other stuff) load the yaml corresponding to the language code
you choose. Please be careful, use the name of the folder as the locale name `string`.
```go
	f := faker.NewFaker("es_ES")
```

You can also change the locale after having your faker initialized
```go
	f := faker.NewFaker("en_GB")
	// do your stuff here with en_GB set
	f.SetLocale("es_ES")
	// Spanish time!
```
## FAQ
 - Is the data real?
   - The data is randomly chosen so no. You may end up with a wrong province/state/zip code combination
- What about gender? Does the library take care of the right combination?
   - Not really. There's a plan to add this capability in the future though.
- What about a CLI or console tool?
   - Is in the roadmap! Check the project regularly
- How can I Help?
   - There's always a missing language or feature. Feel free to open an `issue` or even better, a `pull request` 
   to help this library growing
      
## License
This code is free to use under the terms of the MIT license. See LICENSE.md for more information.