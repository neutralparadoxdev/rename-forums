import { useRouter } from "next/navigation";
import { FC, FormEvent, useState } from "react";

export type NewPostProps = {
    forums: string[],
    rows: number,
    showLogin: () => void
}

export const NewPost: FC<NewPostProps> = ({ forums, rows, showLogin } : NewPostProps) => {
    const router = useRouter();
    if (forums.length === 0) return <></>;

    const [toggle, setToggle] = useState<boolean>(false);
    const [postText, setPostText] = useState<string>("");
    const [forumSelected, setForumSelected] = useState<string>(forums[0]);
    const [title, setTitle] = useState<string>("");

    function expandComponent() {
        const token = localStorage.getItem("session-token") 

        if(token !== null && token !== "") {
            setToggle(val => !val);
            return;
        }

        showLogin();
        return;
    }

    function submit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();
        //console.log(`(title, ${title})(body, ${postText})(forum, ${forumSelected})`)

        const token = localStorage.getItem("session-token")

        if(token !== null) {
            fetch("/api/post", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Bearer-Token": token,
                },
                body: JSON.stringify({
                    title: title,
                    body: postText,
                    forumName: forumSelected
                })
            })
            .then(data => data.json())
            .then(data => {
                router.push(`/f/${forumSelected}/${data.id}`)
            })
            .catch(err => {
                console.log(err)
            })
        }
    }

    return (

        toggle ? 
        (
            <form onSubmit={submit} className="border-2 p-4 mb-2 w-full">
                <div className="mb-2">
                    <h3 className="text-xl font-bold inline mr-2">New Post</h3>
                    <button 
                        onClick={() => setToggle(val => !val)}
                        className="border-2 p-2">Close</button>
                </div>
                <input className="mb-2 p-2 border-2 min-w-full" onChange={e => { setTitle(e.target.value); }} value={title} type="text" alt="title" />
                <textarea onChange={e => { setPostText(e.target.value); }} value={postText} className="border-2 p-1 w-full mb-2 resize-none min-h-max" rows={rows} />
                { forums.length > 1 ? 
                <fieldset>
                    <label htmlFor="forum">Forum</label>
                    <select 
                        onChange={e => { setForumSelected(e.target.value); }} 
                        value={forumSelected} 
                        id="forum" 
                        className="inline mb-2 w-24">{forums.map(x => <option key={x} value={x}>{ x }</option> )}</select> 
                </fieldset>
                : <></> 
                }
                <input className="border-2 p-2 w-24" type="submit" value="Post"/>
            </form>)
        :
        (
            <button 
                onClick={expandComponent}
                className="text-xl font-bold border-2 p-4 mb-2 mt-2 w-full text-left">Create A New Post</button>
        )
    );
}