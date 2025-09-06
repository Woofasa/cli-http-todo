import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="dark bg-background text-text ">
        <div className="flex justify-center">{children}</div>
      </body>
    </html>
  );
}
