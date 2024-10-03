import { CSSProperties } from "react"

export interface chessProps {
    style?: CSSProperties
    className?: string
}

export function App(props: chessProps) {
    return <div className={props.className} style={props.style}>
        Hello Gophers!
    </div>
}