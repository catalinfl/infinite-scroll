import { queryOptions } from "@tanstack/react-query"
import axios from "axios"

export const groupQueries = (id: string) => {
    return queryOptions({
        queryKey: ['user', id],
        queryFn: () => axios.get(`http://localhost:3000/api/${id}`).then(res => res.data)
    })
}

