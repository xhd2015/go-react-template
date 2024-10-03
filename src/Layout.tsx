import { CSSProperties } from "react"
import { Outlet } from "react-router"

export interface LayoutProps {
    style?: CSSProperties
    className?: string
}

export function Layout(props: LayoutProps) {
    return <div className={props.className} style={props.style}>
        <header>
            Welcome To the Go React Template!
        </header>
        <Outlet />
    </div>
}