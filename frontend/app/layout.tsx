import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { ThirdwebProvider } from "thirdweb/react";
import { Toaster } from "@/components/ui/sonner"

const inter = Inter({
    subsets: ["latin"]
});

export const metadata: Metadata = {
    title: "transmuter.xyz",
    description: "Bridge ERC20 tokens with decentralized liquidity pools",
    icons: {
        icon: "/favicon.ico"
    }
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body
                className={inter.className}
            >
                <ThirdwebProvider>
                    <ThemeProvider
                        attribute="class"
                        defaultTheme="dark"
                        disableTransitionOnChange
                    >
                        <main>
                            {children}
                        </main>
                        <Toaster />
                    </ThemeProvider>
                </ThirdwebProvider>
            </body>
        </html>
    );
}
