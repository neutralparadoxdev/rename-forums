'use client';

import { PostStub, PostStubProps } from './PostStub';

import { useState, FC, useEffect } from 'react';

export type PostListProps = {
    posts: PostStubProps[]
}

export const PostList: FC<PostListProps> = ({ posts }: PostListProps) => {
    return (
        <>
        <ul> 
            { posts.map((x) => <PostStub {...x}></PostStub> )}
        </ul>
        </>
    );
}