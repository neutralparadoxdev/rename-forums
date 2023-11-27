'use client';

import { PostStub, PostStubProps } from './PostStub';

import { FC } from 'react';

export type PostListProps = {
    posts: PostStubProps[]
}

export const PostList: FC<PostListProps> = ({ posts }: PostListProps) => {
    return (
        <>
        <ul> 
            { posts.map((x) => <PostStub key={x.id} {...x}></PostStub> )}
        </ul>
        </>
    );
}
