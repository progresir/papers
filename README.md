# Papers


Checking Russian rusdocuments for the checksum.

  - SNILS 
  - INN

### Func:
  - ValidateSnils(snils string) error
  - ValidateInn(inn string) error

### Err:
  - ErrSnils - *Incorrect SNILS*
  -	ErrSnilsCotainInvalidCharacters - *SNILS must consist only of numbers*
  -	ErrINN - *Incorrect INN* 
  