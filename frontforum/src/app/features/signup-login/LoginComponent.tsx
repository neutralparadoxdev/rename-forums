import { FC } from "react";

export type LoginComponentProps = {
    swap: () => void
    close: () => void,
}

export const LoginComponent: FC<LoginComponentProps> = ({ close, swap } : LoginComponentProps) => {
    return (
    <div className="ml-auto mr-auto w-96 bg-sky-500 p-5 mt-44">
        <button className="bg-white rounded-full" onClick={close}>X</button>
        <h1 className="ml-auto mr-auto text-center bold text-2xl">Login</h1>
        <form className="mr-auto ml-auto w-fit">
            <label htmlFor="email" className="block">Email</label>
            <input className="p-1" id="email" type="email" />
            <label htmlFor="password" className="block">Password</label>
            <input id="password" type="password" className="mb-2 p-1" />
            <input className="p-2 bg-white block" type="submit" />
        </form>
        <p className="ml-auto mr-auto w-fit">Do you not have an account? <button className="p-2 bg-white" onClick={swap}>Sign Up</button></p>
    </div>);
}