'use client';
import { useEffect, useState } from 'react';
import { PostList, PostListProps } from '../../features/PostList/PostList';
import { usePathname } from 'next/navigation';
import { LoadingComponent } from '@/app/features/loading/Loading';
import { PostStubProps } from '@/app/features/PostList/PostStub';
import { ErrorComponent } from '../../features/error/Error';

/*
		type PostDTO struct {
			Title      string    `json:"title"`
			AuthorName string    `json:"authorName"`
			CreatedAt  time.Time `json:"createdAt"`
			Id         string    `json:"id"`
		}

		type CompleteForumDTO struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Posts       []PostDTO `json:"posts"`
		}
*/

type PostData = {
  title: string,
  authorName: string,
  createdAt: string,
  id: string,
  subforum: string
}

type ForumData = {
  title: string,
  description: string,
}

function getForumPathParam(pathname : string) {
  const pathSplit = pathname.split('/')
  const urlParam = pathSplit[pathSplit.length-1]
  return urlParam
}

export default function ForumPage() {
  const pathname = usePathname()
  const pathParam = getForumPathParam(pathname)

  const [isLoading, setIsLoading] = useState<boolean>(true);

  const [forumData, setForumData] = useState<ForumData| null>(null);
  const [posts, setPosts] = useState<PostData[] | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("/api/forum/" + pathParam)
    .then(data => {
      return data.json()
    })
    .then((data : any) => {
      setPosts(data.posts.map((post: PostData) => {post.subforum = pathname; return post}));
      setForumData({
        title: data.title,
        description: data.description,
      });
    })
    .catch(err => {
      setError(err);
    })
    .finally(() => {
      setIsLoading(false);
    }
    )
  })

  const page = (
    <main className="border-box p-2 min-w-full">
      <h1 className="text-2xl min-w-full border-b-2">{forumData !== null ? forumData.title : "Lorem Ipsum"}</h1>
      <PostList posts={posts !== null ? posts : []} />
    </main>
  )

  return (isLoading ? <LoadingComponent /> : (error !== null || true ? page : <ErrorComponent msg={error !== null ? "Some Error" : ""} />));
}
