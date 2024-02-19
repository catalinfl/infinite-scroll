import { useQuery } from "@tanstack/react-query"
import axios from "axios"
import { groupQueries } from "./groupQueries"

const fetchID = async (id: string) => {
    const res = await axios.get(`http://localhost:3000/api/${id}`)
    return res.data
}



export const FileIDQuery = () => {
  
  return (
    <div className="application">
        <UserID userid="3" />
        <UserID userid="3" />
    </div>
  )

} 

function UserID({ userid }: { userid: string }) {
    const result = useQuery(groupQueries(userid))
    if (result.isLoading) return <div>Loading...</div>

    if (result.isError) return <div>Error: {result.error.message}</div>

    return <div>
        <p> {result.data.name} </p>
        <p> {result.data.email} </p>
        </div>
        
    
}
