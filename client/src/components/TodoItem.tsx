import { BASE_URL } from "../App";
import { Badge, Box, CircularProgress, Flex, Spinner, Text } from "@chakra-ui/react";
import { QueryClient, useMutation, useQueryClient } from "@tanstack/react-query";
import React from "react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";

const TodoItem = ({ todo }: { todo: any }) => {

    const queryClient = useQueryClient();
    const {mutate:updateTodo,isPending: isUpdating} = useMutation ({
        mutationKey:["updateTodo"],
        mutationFn: async() => {
            if (todo.completed) return alert("Todo already Completed")
                try {
                    const res = await fetch(BASE_URL + `/todos/${todo._id}`,{
                        method: "PATCH"
                    })

                    const data = await res.json()

                    if (!res.ok) { 
                    throw new Error(data.error || "Something Went wrong")
                    }

                return data
                } catch (error) {
                    console.log(error)
                    
                }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]});

        }
    });
    const {mutate:deleteTodo,isPending: isdeleting} = useMutation ({
        mutationKey:["deleteTodo"],
        mutationFn: async() => {
                try {
                    const res = await fetch(BASE_URL + `/todos/${todo._id}`,{
                        method: "delete"
                    })

                    const data = await res.json()

                    if (!res.ok) { 
                    throw new Error(data.error || "Something Went wrong")
                    }

                return data
                } catch (error) {
                    console.log(error)
                    
                }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]});

        }
    });

	return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.600"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={todo.completed ? "green.200" : "blue.500"}
					textDecoration={todo.completed ? "line-through" : "none"}
				>
					{todo.body}
				</Text>
				{todo.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!todo.completed && (
					<Badge ml='1' colorScheme='blue'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"} onClick={() => updateTodo()}>
                    {!isUpdating && <FaCheckCircle size={20} />}
                    {isUpdating && <Spinner size={"small"} />}
				</Box>
				<Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
                    {!isdeleting && <MdDelete size={25} />}
                    {isdeleting && <Spinner size={"small"} />}
				</Box>
			</Flex>
		</Flex>
	);
};
export default TodoItem;
