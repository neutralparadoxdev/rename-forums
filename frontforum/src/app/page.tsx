import Image from 'next/image'

import { PostList } from './features/post-list/PostList';

export default function Home() {
  return (
    <main className="border-box p-2">
      <h1 className="text-2xl">Welcome home</h1>
      <PostList for="/home"></PostList>
    </main>
  )
}
