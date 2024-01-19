package prueba

async function controller2(req, res) {
	const prueba = "prueba"
const conn = await getConnection()
conn.query("SELECT * FROM Post WHERE id = ?", req.params.id);
// comentario
}