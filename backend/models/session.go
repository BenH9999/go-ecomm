package models

import(
    "sync"
)

var SessionStore = make(map[string]string)
var sessionMutex sync.RWMutex

func SetSession(token, username string){
    sessionMutex.Lock()
    defer sessionMutex.Unlock()
    SessionStore[token] = username
}

func GetSession(token string) (string, bool){
    sessionMutex.RLock()
    defer sessionMutex.RUnlock()
    username, ok := SessionStore[token]
    return username, ok
}

func DeleteSession(token string){
    sessionMutex.Lock()
    defer sessionMutex.Unlock()
    delete(SessionStore, token)
}
