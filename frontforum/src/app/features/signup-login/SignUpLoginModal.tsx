import { FC } from "react";
import { LoginComponent } from "./LoginComponent";
import { SignUpComponent } from "./SignUpComponent";

export enum SignUpLoginModalPurpose {
    Login,
    SignUp,
} 

export type SignUpLoginModal = {
    purpose: SignUpLoginModalPurpose,
    changePurpose: (purpose: SignUpLoginModalPurpose) => void,
    close: () => void
}


export const SignUpLoginModal: FC<SignUpLoginModal> = ({purpose, close, changePurpose} : SignUpLoginModal) => {

    let internal;
    switch(purpose) {
        case SignUpLoginModalPurpose.Login:
            internal = <LoginComponent close={close} swap={() => changePurpose(SignUpLoginModalPurpose.SignUp)}/>
            break;
        case SignUpLoginModalPurpose.SignUp:
            internal = <SignUpComponent close={close} swap={() => changePurpose(SignUpLoginModalPurpose.Login)} />
            break;
    }

    let ret = (
    <>
    <div className="absolute h-screen border-2 min-w-fit w-screen bg-sky-300/50">
        { internal }
    </div>
    </>)

    return ret
}
