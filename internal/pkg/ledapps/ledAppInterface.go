package ledapps

type LedAppInterface interface {
    Setup()
    Loop()
    Cleanup()
}
