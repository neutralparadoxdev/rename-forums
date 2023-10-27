import { FC } from "react";

export type SignUpComponentProps = {
    swap: () => void,
    close: () => void,
}

export const SignUpComponent: FC<SignUpComponentProps> = ({close, swap} : SignUpComponentProps) => {
    return (
    <div className="ml-auto mr-auto w-96 bg-sky-500 p-5 mt-44">
        <button className="bg-white rounded-full" onClick={close}>X</button>
        <h1 className="ml-auto mr-auto text-center bold text-2xl">Create Account</h1>
        <form className="ml-auto mr-auto w-fit">
            <label htmlFor="email" className="block">Email</label>
            <input id="email" type="email" className="p-1" required/>
            <label htmlFor="username" className="block">Username</label>
            <input id="username" type="text" className="p-1" required/>
            <label htmlFor="password" className="block">Password</label>
            <input id="password" type="password" className="p-1" required/>
            <label htmlFor="password2" className="block">Password Confirmation</label>
            <input id="password2" className="block p-1" type="password" required/>
            <label htmlFor="eula" className="">Do you accept <a className="underline mr-2" href="">eula</a></label>
            <input id="eula" type="checkbox" required/>
            <input className="p-2 bg-white block" type="submit" />
        </form>
        <p className="ml-auto mr-auto w-fit">Already have an account? <button onClick={swap} className="p-2 bg-white">Login</button></p>
    </div>);

}