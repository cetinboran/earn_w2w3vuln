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
                    throw new Error(`Backend hata: ${response.status} ${response.statusText}`);
                }

                const data = await response.json();

                setNfts(data.data);
            } catch (error) {
                console.log("Fetch sırasında hata oluştu:", error);
            } finally {
                setLoading(false);
            }
        };

        fetchNFTs();
    }, [user.id]);

    const handleMint = async (nftId) => {
        if (!user || !user.publicKey) {
            alert("Kullanıcı bilgisi eksik.");
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

            console.log(user.id, nftId, user.publicKey)

            const result = await response.json();
            console.log(result, "Result")

            if (!response.ok) {
                throw new Error(result.message || "Mint işlemi başarısız.");
            }

            setMinted([...minted, nftId]);
            alert(`NFT ${nftId} başarıyla mintlendi!`);
        } catch (error) {
            console.log("Mint işlemi sırasında hata:", error);
            alert("Mint sırasında bir hata oluştu.");
        }
    };


    return (
        <div style={{ padding: "20px", fontFamily: "Arial" }}>
            <h2>Hoş geldin, {user.username}!</h2>
            <p>Kullanıcı ID: {user.id}</p>

            <button
                onClick={() => {
                    localStorage.removeItem('user');
                    window.location.reload();
                }}
                style={{ margin: "20px 0" }}
            >
                Çıkış Yap
            </button>

            <h3>Mint Edilebilir NFT'ler</h3>

            {loading ? (
                <p>Yükleniyor...</p>
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
                                        disabled={minted.includes(nft.id)}
                                        style={{ marginLeft: '20px', padding: '6px 12px' }}
                                    >
                                        {minted.includes(nft.id) ? "Mintlendi" : "Mint Et"}
                                    </button>
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p>Hiç NFT bulunamadı.</p>
                    )}
                </>
            )}
        </div>
    );
};

export default Dashboard;
