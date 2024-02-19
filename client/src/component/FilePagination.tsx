import { keepPreviousData, useQuery } from '@tanstack/react-query'
import axios from 'axios'
import React, { useState } from 'react'

const FilePagination = () => {

    const [page, setPage] = useState<number>(0)

    const fetchProject = async (page = 0) => {
        return axios.get(`http://localhost:3000/api?page=${page}`)
            .then((res) => {
                return res.data
            })
    }

    console.log(page)

    const { isPending, isError, error, data, isFetching, isPlaceholderData } =
    useQuery({
      queryKey: ['projects', page],
      queryFn: () => fetchProject(page),
      placeholderData: keepPreviousData,
    })



    return (
        <div>
            {isPending ? (
                <div> Loading </div>
            ) : isError ? (
                <div> Error: {error.message} </div>
            ) : (
                <div>
                    {data?.map((user: any) => {
                        return (
                            <div key={user.id}>
                                <p> {user.name} </p>
                                <p> {user.email} </p>
                            </div>
                        )
                    })}
                </div>
            )}
            <button onClick={() => setPage(Math.max(page - 1, 0))} disabled={page === 0}> Previous </button>
            <button onClick={() => { 
                if (data?.length === 2) {
                    setPage(page + 1)
                }
                }}
                disabled={data?.length < 2}
                > Next </button>
            {isFetching ? <span> Loading...</span> : null}
        </div>
      )
    }


export default FilePagination