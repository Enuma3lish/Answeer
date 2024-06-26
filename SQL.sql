1.SELECT bnb_id,bnb_name,SUM(amount) as may_amount FROM orders
	Where currency = “ TWD “
		AND create_at >= ‘2023-05–01’
		AND create_at< ‘2023-06-01
	GROUP BY bnb_id,bnb_name
	ORDER BY may_amount DESC
	LIMIT10;
2.1.對於比較大的表可以進行分區查詢．
  2.檢查確保表結構合理。若 bnb_name 存在於另一個表中（例如 bnb_info），則使用聯接而非在 orders 表中直接存儲名稱。這樣可以減少多餘的query，提高查詢效率。

