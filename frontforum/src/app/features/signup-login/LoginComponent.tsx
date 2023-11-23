import { SetSessionToken, SetUsername } from "@/app/services/SessionManager/session";
import { useRouter } from "next/navigation";
import { EventHandler, FC, FormEvent, FormEventHandler, useState } from "react";

export type LoginComponentProps = {
    swap: () => void
    close: () => void,
    setAuthToken: (token: string) => void
}

export const LoginComponent: FC<LoginComponentProps> = ({ close, swap, setAuthToken } : LoginComponentProps) => {
    const [password, setPassword] = useState<string>("");
    const [username, setUsername] = useState<string>("");
    const router = useRouter();

    function submit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
        console.log(`(username:${username})(password:${password})`)
        console.log(JSON.stringify({ 'username' : username, 'password' : password}))
        fetch("/api/session/new", {
            headers: {
                "Content-Type" : "application/json",
            },
            method: 'POST',
            body: JSON.stringify({ 'username' : username, 'password' : password}),
            
        })
        .then(data => {
            if(data.ok) {
                return data.json()
            }
            throw Error("response is not ok")
        })
        .then(data => {
            SetSessionToken(data.token)
            SetUsername(data.username)
            setAuthToken(data.token)
            router.refresh()
        })
        .catch(err => {
            console.log(err)
        })
    }

    return (
    <div className="ml-auto mr-auto w-96 bg-sky-500 p-5 mt-44">
        <button className="bg-white rounded-full" onClick={close}>X</button>
        <h1 className="ml-auto mr-auto text-center bold text-2xl">Login</h1>
        <form className="mr-auto ml-auto w-fit" onSubmit={submit}>
            <label htmlFor="username" className="block">Username</label>
            <input onChange={(e) => {setUsername(e.target.value)}} value={username} className="p-1" id="username" type="text" />
            <label htmlFor="password" className="block">Password</label>
            <input onChange={(e) => {setPassword(e.target.value)}} value={password} id="password" type="password" className="mb-2 p-1" />
            <input className="p-2 bg-white block" type="submit" />
        </form>
        <p className="ml-auto mr-auto w-fit">Do you not have an account? <button className="p-2 bg-white" onClick={swap}>Sign Up</button></p>
    </div>);
}