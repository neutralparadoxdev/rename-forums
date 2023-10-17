'use client';

export type PostStubProps = {
    title: string
    authorName: string
    id: string
    subforum: string
}

function genLink(): string {
    return ""
}

export const PostStub: React.FC<PostStubProps> = ({ title, user, subforum, id }: PostStubProps) => {
    return (
        <li key={(user + subforum + id)}
            className="p-1 border-2 min-w-max grid grid-rows-4 grid-cols-[3%_98%]">
            <a className="grid-row-1 col-start-1 col-end-1 row-span-2" href={"hover:border-2 hover:text-red-500"}>👍</a>
            <a className="grid-row-3 col-start-1 col-end-1 row-span-2" href={"hover:border-2 hover:text-red-500"}>👎</a>
            <div className="mt-[auto] mb-[auto] row-start-1 row-end-5 row-span-4 col-start-2 col-end-3 align-right w-max-full">
                <span className=" font-bold text-2xl hover:border-2 hover:text-red-500 mt-auto mb-auto" ><a  href={genLink()}>{title}</a></span>
                <span>by {user}</span>
            </div>
        </li>
        );
}