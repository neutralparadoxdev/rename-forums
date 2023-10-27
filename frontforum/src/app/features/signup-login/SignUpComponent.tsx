import { useRouter } from "next/navigation";
import { FC, FormEvent, useState } from "react";

export type SignUpComponentProps = {
    swap: () => void,
    close: () => void,
}

export const SignUpComponent: FC<SignUpComponentProps> = ({close, swap} : SignUpComponentProps) => {

    const router = useRouter();

    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [password2, setPassword2] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [eula, setEula] = useState<boolean>(false);

    function submit(e: FormEvent<HTMLFormElement>){
        e.preventDefault();

        if(password != password2) return;

        if(!eula) return

        console.log(JSON.stringify({ 'username' : username, 'email' : email, 'password' : password, 'eula' : eula}))

        fetch("/api/user", {
            headers: {
                "Content-Type" : "application/json",
            },
            method: 'POST',
            body: JSON.stringify({ 'username' : username, 'email' : email, 'password' : password, 'eula' : eula}),
        })
        .then(data => {
            console.log(data.status)
            if(data.status == 201) {
                swap();
                router.refresh();
                return;
            }
            throw Error("response is not ok")
        })
        .catch(err => {
            console.log(err)
        })
    }

    return (
    <div className="ml-auto mr-auto w-96 bg-sky-500 p-5 mt-44">
        <button className="bg-white rounded-full" onClick={close}>X</button>
        <h1 className="ml-auto mr-auto text-center bold text-2xl">Create Account</h1>
        <form className="ml-auto mr-auto w-fit" onSubmit={submit}>
            <label htmlFor="email" className="block">Email</label>
            <input onChange={(e) => setEmail(e.target.value)} value={email} id="email" type="email" className="p-1" required/>
            <label htmlFor="username" className="block">Username</label>
            <input onChange={(e) => setUsername(e.target.value)} value={username} id="username" type="text" className="p-1" required/>
            <label htmlFor="password" className="block">Password</label>
            <input onChange={(e) => setPassword(e.target.value)} value={password} id="password" type="password" className="p-1" required/>
            <label htmlFor="password2" className="block">Password Confirmation</label>
            <input onChange={(e) => setPassword2(e.target.value)} value={password2} id="password2" className="block p-1" type="password" required/>
            <label htmlFor="eula" className="">Do you accept <a className="underline mr-2" href="">eula</a></label>
            <input onChange={(e) => setEula((v) => !v)} checked={eula} id="eula" type="checkbox" required/>
            <input className="p-2 bg-white block" type="submit" />
        </form>
        <p className="ml-auto mr-auto w-fit">Already have an account? <button onClick={swap} className="p-2 bg-white">Login</button></p>
    </div>);

}