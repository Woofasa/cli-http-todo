import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className=" ">
        <div className="flex justify-center">{children}</div>
      </body>
    </html>
  );
}
