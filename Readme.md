## ed25519 base Key Generation

How to get key value pair?

```
make run 
```

You will see on your log:

```
Public Key: mkjr+2Ihki7Sd/M6t0FMnPB4kaYgob3ZIIUy9MUyAUw=
Private Key: 28sE9P5CxDSxDuvJvfHSHL4/rVxNkYaw+C6gdKl8CSOaSOv7YiGSLtJ38zq3QUyc8HiRpiChvdkghTL0xTIBTA==
```

These Key Value Pairs are base64 encoded.
To load an ed25519
```go
    pubKeyBytes, err := base64.StdEncoding.DecodeString(PublicKey)
    if err != nil {
        return fmt.Errorf("pubKey decoding failure: %s", err.Error())
    }
    privKeyBytes, err := base64.StdEncoding.DecodeString(PrivateKey)
    if err != nil {
        return fmt.Errorf("pubKey decoding failure: %s", err.Error())
    }
    publicKey := ed25519.PubKey{Key: pubKeyBytes}
    privateKey := ed25519.PrivKey{Key: privKeyBytes}
```