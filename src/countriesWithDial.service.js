import datas from './datas.json';

export const infos = async (req, res) => {
    try {
        return res.status(200).json(datas)
    } catch(error) {
        return res.status(500).json(error)
    }
}