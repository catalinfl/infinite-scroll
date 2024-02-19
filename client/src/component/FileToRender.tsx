import { useQuery } from "@tanstack/react-query"
import axios from "axios"

export const FileToRender = () => {
  const { isPending, error, data } = useQuery({
    queryKey: ['repoData'],
    queryFn: () =>
      axios.get('http://localhost:3000/api').then((res: any) => res.data),
  })


  if (isPending) return <div>Loading...</div>

  if (error) return <div>Error: {error.message}</div>

  console.log(data)
  
  return (
    <div className="application">
      {data?.map((user: any) => {
        return (
          <div className="user" key={user.id}>
            <p> Name: {user.name} </p>
            <p> User: {user.email} </p>
          </div>
        )
      })}
    </div>
  )

} 