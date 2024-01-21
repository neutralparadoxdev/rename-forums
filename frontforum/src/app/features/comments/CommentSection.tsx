import { FC } from 'react';

import { CommentProps, Comment } from './Comment';

export type { CommentProps } from './Comment';

export type CommentSectionProps = {
	comments: CommentProps[]
}

export const CommentSection: FC<CommentSectionProps> = ({ comments }:CommentSectionProps) => {
	return (
		<div className="border-l-4 pl-4 mt-2 height-full width-full">
		{ 
			comments.map((comment: CommentProps) => {
				return <Comment {...comment} /> 
			})
		} 
		</div>
	);
}
