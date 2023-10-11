import Image from 'next/image'

import { PostList } from '../features/PostList/PostList';

export default function Home() {
  return (
    <main className="border-box p-2 min-w-full">
      <h1 className="text-2xl min-w-full border-b-2">Welcome home</h1>
      <PostList for="/home"></PostList>
    </main>
  )
}
