module Stack = struct
  type 'a t = 'a list
  
  let empty () = []
  
  let is_empty stack = stack = []
  
  let push x stack = x :: stack
  
  let pop stack =
    match stack with
    | [] -> None
    | x :: xs -> Some (x, xs)
  
  let peek stack =
    match stack with
    | [] -> None
    | x :: _ -> Some x
end

let () =
  let stack = Stack.empty () in
  let stack = Stack.push 1 stack in
  let stack = Stack.push 2 stack in
  let stack = Stack.push 3 stack in
  match Stack.pop stack with
  | None -> print_endline "Stack is empty"
  | Some (x, _) -> print_endline ("Popped: " ^ string_of_int x)
