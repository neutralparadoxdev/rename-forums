'use client';
import { useEffect, useState } from 'react';
import { PostList } from '../../features/PostList/PostList';
import { usePathname } from 'next/navigation';
import { LoadingComponent } from '@/app/features/loading/Loading';
import { ErrorComponent } from '../../features/error/Error';
import { Header } from '@/app/features/header/header';

import { NewPost } from '../../features/NewPost/NewPost';
import { SignUpLoginModalPurpose } from '@/app/features/signup-login/SignUpLoginModal';
import { GetSessionToken } from '@/app/services/SessionManager/session';

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

  const [loginSignUpState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);


  useEffect(() => {

    const sessionToken = GetSessionToken()

    if (sessionToken === null|| sessionToken === "") {
      fetch("/api/forum/" + pathParam)
      .then(data => {
        if(data.status === 404 || data.status === 401) {
          setError("Not Found")
          throw Error("Not Found")
        }
        return data.json()
      })
      .then((data : any) => {
        setPosts(data.posts.map((post: PostData) => {post.subforum = pathParam; return post}));
        setForumData({
          title: data.title,
          description: data.description,
        });
      })
      .catch(err => {
		console.log(err)
        setError("Not Found");
      })
      .finally(() => {
        setIsLoading(false);
      }
      )
      return
    }
    else {
      fetch("/api/forum/" + pathParam, {
        headers: {
          "Bearer-Token" : sessionToken
        }
      })
      .then(data => {
        return data.json()
      })
      .then((data : any) => {
        setPosts(data.posts.map((post: PostData) => {post.subforum = pathParam; return post}));
        setForumData({
          title: data.title,
          description: data.description,
        });
      })
      .catch(err => {
		console.log(err)
        setError("Not Found");
      })
      .finally(() => {
        setIsLoading(false);
      }
      )
      return


    }
  }, [])

  const page = (
    <>
    <Header 
      title={forumData !== null ? forumData.title : "Lorem Ipsum"} 
      link={null} 
      setLoginSignUpState={x => { setLoginSignUpState(x);}}
      loginSignUpState={loginSignUpState}
      />
    <main className="border-box p-2 min-w-full">
      <NewPost 
        rows={8} 
        forums={[forumData !== null ? forumData.title : "lorem"]}
        showLogin={() => setLoginSignUpState(SignUpLoginModalPurpose.Login)}/>
      <PostList posts={posts !== null ? posts : []} />
    </main>
    </>
  )

  return (isLoading ? <LoadingComponent /> : (error === null ? page : <ErrorComponent msg={error} />));
}
