import { windowChecker } from '../WindowChecker/windowChecker';

const SESSION_TOKEN = 'session-token';
const USERNAME = 'session-token-username';

export function GetSessionToken(): string | null {
	if(!windowChecker()) return null;
    return localStorage.getItem(SESSION_TOKEN);
}

export function SessionTokenExist(): boolean {
	if(!windowChecker()) return false;
    let token = localStorage.getItem(SESSION_TOKEN)

    return token !== null && token !== ""
}

export function SetSessionToken(token: string) {
	if(!windowChecker()) return;
	localStorage.setItem(SESSION_TOKEN, token);
}

export function SetUsername(username: string) {
     localStorage.setItem(USERNAME, username);
}

export function DeleteSessionToken() {
	if(!windowChecker()) return;
    localStorage.setItem(SESSION_TOKEN, '')
    localStorage.setItem(USERNAME, '')
}

export function GetUsername(): string | null {
	if(!windowChecker()) return null;
    if(SessionTokenExist()) {
        return localStorage.getItem(USERNAME)
    }
    return null
}
