'use client';

import { PostStub, PostStubProps } from './PostStub';

import { useState, FC } from 'react';

type PostListProps = {
    for: string
}

const EXAMPLE_POSTS: PostStubProps[] = [
    { id: "none", subforum: "home", title: "Hello World", user: "Some user" },
    { id: "none1", subforum: "home", title: "Hello World 2", user: "Some user" }
]

export const PostList: FC<PostListProps> = ({}: PostListProps) => {

    const [data, setData] = useState<PostListProps[]>([]);

    return (
        <>
        <ul> 
            { EXAMPLE_POSTS.map((x) => <PostStub {...x}></PostStub> )}
        </ul>
        </>
    );
}