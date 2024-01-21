import { FC } from 'react';

export type CommentProps = {
	text: string,
	username: string,
	userid: string,
	comments: CommentProps[],
}

export const Comment: FC<CommentProps> = ({ text, comments,  username } : CommentProps) => {
	return (
		<div className="border-l-4 pl-4 mt-2 height-full width-full border-b-2 last:border-b-0">
		<p>{ text }</p>
		<p><a href="">{username}</a></p>
		<div>
			<button className="border-2 border-blue-600 text-blue-600 w-[18]">Reply</button>
			<button>👍</button>
			<button>👎</button>
		</div>
		<div className="">
		{ 
			comments.map(comment => {
				return <Comment {...comment} />
			})
		}
		</div>
		</div>
	);
}
