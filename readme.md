# btc puzzle solver

https://btcpuzzle.info/puzzlelist


## TODO

 - perfomance metrics
 - telegram notify
 - threads for 100% cpu usage
 - GPU usage
 - electrum sidecar
   - check balance
   - make transaction if found key
 - tor sidecar for privacy
 - k8s manifests
 - split app via nsq (separate workers for key gen and so on)

## 65-72_9byte

Simple mvp go app for puzzles 65 - 72 exclude solved 65,66,70

    65 Solved
    66 Solved
    67 1BY8GQbnueYofwSuFAT3USAhGjPrkxDdW9 6.70010696 BTC
    68 1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ 6.80005314 BTC
    69 19vkiEajfhuZ8bs8Zu2jgmC6oqZbWqhxhG 6.90013061 BTC
    70 Solved
    71 1PWo3JeB9jrGwfHDNpdGK54CRas7fsVzXU 7.10003770 BTC
    72 1JTK7s9YVYywfm5XUH7RNhHJH1LshCaRFR 7.20003779 BTC

    34.8 BTC total ($3.5M)

## 73-80_10byte

### 73 - Unsolved

    Key Range                           Reward
    (2^72)...(2^73)                     7.30004377 BTC

    Start Point (Decimal)               End Point (Decimal)
    4722366482869645213696 = (2^72)     9444732965739290427391 = (2^73)

    Start HEX                           End HEX
    1000000000000000000                 1ffffffffffffffffff

    Total Keys                          Target
    4722366482869645213696              12VVRNPi4SJqUTsp6FmqDqY5sGosDtysn4

### 74 - Unsolved

    Key Range                           Reward
    (2^73)...(2^74)                     7.40004377 BTC

    Start Point (Decimal)               End Point (Decimal)
    4722366482869645213696 = (2^73)     18889465931478580854783 = (2^74)

    Start HEX                           End HEX
    2000000000000000000                 3ffffffffffffffffff

    Total Keys                          Target
    9444732965739290427392              1FWGcVDK3JGzCC3WtkYetULPszMaK2Jksv

### 75 - Solved

    Key Range                           Reward
    (2^74)...(2^75)                     0 BTC

    Start Point (Decimal)               End Point (Decimal)
    18889465931478580854784 = (2^74)    37778931862957161709567 = (2^75)

    Start HEX                           End HEX
    4000000000000000000                 7ffffffffffffffffff

    Total Keys                          Target
    18889465931478580854784             1J36UjUByGroXcCvmj13U6uwaVv9caEeAt

    🔑Private Key
    0x4c5ce114686a1336e07

### 76 - Unsolved

    Key Range                           Reward
    (2^75)...(2^76)                     7.6 BTC

    Start Point (Decimal)               End Point (Decimal)
    37778931862957161709568 = (2^75)    75557863725914323419135 = (2^76)

    Start HEX                           End HEX
    8000000000000000000                 fffffffffffffffffff

    Total Keys                          Target
    37778931862957161709568             1DJh2eHFYQfACPmrvpyWc8MSTYKh7w9eRF

### 77 - Unsolved

    Key Range                           Reward
    (2^76)...(2^77)                     7.70002426 BTC

    Start Point (Decimal)               End Point (Decimal)
    75557863725914323419136 = (2^76)    151115727451828646838271 = (2^77)

    Start HEX                           End HEX
    10000000000000000000                1fffffffffffffffffff

    Total Keys                          Target
    75557863725914323419136             1Bxk4CQdqL9p22JEtDfdXMsng1XacifUtE

### 78 - Unsolved

    Key Range                           Reward
    (2^77)...(2^78)                     7.8 BTC

    Start Point (Decimal)               End Point (Decimal)
    151115727451828646838272 = (2^77)   302231454903657293676543 = (2^78)

    Start HEX                           End HEX
    20000000000000000000                3fffffffffffffffffff

    Total Keys                          Target
    151115727451828646838272            15qF6X51huDjqTmF9BJgxXdt1xcj46Jmhb

### 79 - Unsolved

    Key Range                           Reward
    (2^78)...(2^79)                     7.9 BTC

    Start Point (Decimal)               End Point (Decimal)
    302231454903657293676544 = (2^78)   604462909807314587353087 = (2^79)

    Start HEX                           End HEX
    40000000000000000000                7fffffffffffffffffff

    Total Keys                          Target
    302231454903657293676544            1ARk8HWJMn8js8tQmGUJeQHjSE7KRkn2t8

### 80 - Solved

    Key Range                           Reward
    (2^79)...(2^80)                     0 BTC

    Start Point (Decimal)               End Point (Decimal)
    604462909807314587353088 = (2^79)   1208925819614629174706175 = (2^80)

    Start HEX                           End HEX
    80000000000000000000                ffffffffffffffffffff

    Total Keys                          Target
    604462909807314587353088            1BCf6rHUW6m3iH2ptsvnjgLruAiPQQepLe

    🔑Private Key
    0xea1a5c66dcc11b5ad180


## 3byte_key

For testing purposes

Solved puzzles 17-24

    puzzle, key range, key, addr
    17 010000-01ffff 0x01764f 1HduPEXZRdG26SUT5Yk83mLkPyjnZuJ7Bm
    18 020000-03ffff 0x03080d 1GnNTmTVLZiqQfLbAdp9DVdicEnB5GoERE
    19 040000-07ffff 0x05749f 1NWmZRpHH4XSPwsW6dsS3nrNWfL1yrJj4w
    20 080000-0fffff 0x0d2c55 1HsMJxNiV7TLxmoF6uJNkydxPFDog4NQum
    21 100000-1fffff 0x1ba534 14oFNXucftsHiUMY8uctg6N487riuyXs4h
    22 200000-3fffff 0x2de40f 1CfZWK1QTQE3eS9qn61dQjV89KDjZzfNcv
    23 400000-7fffff 0x556e52 1L2GM8eE7mJWLdo3HZS6su1832NX2txaac
    24 800000-ffffff 0xdc2a04 1rSnXMr63jdCuegJFuidJqWxUPV7AtUf7
