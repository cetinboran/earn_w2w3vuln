import React, { useEffect, useState } from 'react';

const Dashboard = ({ user }) => {
    const [nfts, setNfts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [minted, setMinted] = useState([]);

    useEffect(() => {
        const fetchNFTs = async () => {
            try {
                const response = await fetch(`http://localhost/api/user/getNfts/${user.id}`, {
                    method: 'GET',
                });

                if (!response.ok) {
                    throw new Error(`Backend Error: ${response.status} ${response.statusText}`);
                }

                const data = await response.json();

                setNfts(data.data);
            } catch (error) {
                console.log("Error occurred during fetch:", error);
            } finally {
                setLoading(false);
            }
        };

        fetchNFTs();
    }, [user.id]);

    const handleMint = async (nftId) => {
        if (!user || !user.publicKey) {
            alert("User information is missing.");
            return;
        }

        try {
            const response = await fetch("http://localhost/api/nft/mint/demo", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    userID: user.id,
                    nftID: nftId,
                    publicKey: user.publicKey,
                }),
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(result.message || "Mint operation failed.");
            }

            setMinted([...minted, nftId]);
            alert(`NFT ${nftId} successfully minted!`);
        } catch (error) {
            alert("An error occurred during Mint.");
        }
    };


    return (
        <div style={{ padding: "20px", fontFamily: "Arial" }}>
            <h2>Welcome, {user.username}!</h2>
            <p>User ID: {user.id}</p>

            <button
                onClick={() => {
                    localStorage.removeItem('user');
                    window.location.reload();
                }}
                style={{ margin: "20px 0" }}
            >
                Logout
            </button>

            <h3>Can be Minted</h3>

            {loading ? (
                <p>Loading...</p>
            ) : (
                <>
                    {nfts.length > 0 ? (
                        <ul style={{ listStyleType: 'none', padding: 0 }}>
                            {nfts.map((nft) => (
                                <li key={nft.id} style={{
                                    marginBottom: '20px',
                                    border: '1px solid #ccc',
                                    borderRadius: '8px',
                                    padding: '10px',
                                    display: 'flex',
                                    alignItems: 'center'
                                }}>
                                    <img
                                        src="/images.png"
                                        alt="NFT"
                                        style={{ width: '80px', height: '80px', marginRight: '15px', borderRadius: '8px' }}
                                    />
                                    <div style={{ flexGrow: 1 }}>
                                        <strong>Name:</strong> {nft.name} <br />
                                        <strong>Symbol:</strong> {nft.symbol} <br />
                                        <strong>URI:</strong> <a href={nft.uri} target="_blank" rel="noopener noreferrer">{nft.uri}</a> <br />
                                        <strong>Seller Fee:</strong> {nft.seller_fee}
                                    </div>
                                    <button
                                        onClick={() => handleMint(nft.id)}
                                        disabled={nft.isMinted}
                                        style={{ marginLeft: '20px', padding: '6px 12px' }}
                                    >
                                        {minted.includes(nft.id) ? "Minted" : "Mint NFT"}
                                    </button>
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p>NFT Not Found.</p>
                    )}
                </>
            )}
        </div>
    );
};

export default Dashboard;
