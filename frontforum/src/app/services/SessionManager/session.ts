const SESSION_TOKEN = 'session-token';
const USERNAME = 'session-token-username';

export function GetSessionToken(): string | null {
    return localStorage.getItem(SESSION_TOKEN);
}

export function SessionTokenExist(): boolean {
    let token = localStorage.getItem(SESSION_TOKEN)

    return token !== null && token !== ""
}

export function SetSessionToken(token: string) {
     localStorage.setItem(SESSION_TOKEN, token);
}

export function SetUsername(username: string) {
     localStorage.setItem(USERNAME, username);
}

export function DeleteSessionToken() {
    localStorage.setItem(SESSION_TOKEN, '')
    localStorage.setItem(USERNAME, '')
}

export function GetUsername(): string | null {
    if(SessionTokenExist()) {
        return localStorage.getItem(USERNAME)
    }
    return null
}