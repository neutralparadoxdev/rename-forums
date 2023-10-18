'use client';

export type PostStubProps = {
    title: string
    authorName: string
    id: string
    forum: string
}

function genLink(): string {
    return ""
}

export const PostStub: React.FC<PostStubProps> = ({ title, authorName, forum, id }: PostStubProps) => {
    return (
        <li key={(id)}
            className="p-1 border-2 min-w-max grid grid-rows-4 grid-cols-[3%_98%]">
            <a className="grid-row-1 col-start-1 col-end-1 row-span-2 hover:border-2 hover:text-red-500" href={""}>👍</a>
            <a className="grid-row-3 col-start-1 col-end-1 row-span-2 hover:border-2 hover:text-red-500" href={""}>👎</a>
            <div className="mt-[auto] mb-[auto] row-start-1 row-end-5 row-span-4 col-start-2 col-end-3 align-right w-max-full">
                <span className=" font-bold text-2xl hover:border-2 hover:text-red-500 mt-auto mb-auto" ><a  href={"/f/" + forum + "/" + id}>{title}</a></span>
                <span>by {authorName}</span>
            </div>
        </li>
        );
}