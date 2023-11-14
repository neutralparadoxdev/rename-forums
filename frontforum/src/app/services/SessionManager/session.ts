export function GetSessionToken(): string | null {
    return localStorage.getItem('session-token');
}

export function SessionTokenExist(): boolean {
    let token = localStorage.getItem('session-token');

    return token !== null && token !== ""
}