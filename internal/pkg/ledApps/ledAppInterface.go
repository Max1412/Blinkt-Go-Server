package ledApps

type LedAppInterface interface {
    Setup()
    Loop()
    Cleanup()
}
