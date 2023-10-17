import { FC } from "react"

export type ErrorProps = {
    msg: string
}

export const ErrorComponent: FC<ErrorProps> = ({ msg } : ErrorProps) => {
    return (
        <>
        <h1>Error: {msg}</h1>
        </>
    );
}