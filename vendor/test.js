function checkCashRegister(price, cash, cid) {
  let status, change;
  change = cash - price;

  function findAmount(status, changeDue, changed) {
    if (status == 'open' && changeDue > 0) {
      return findAmount(status, changeDue)
    } else {
       const totalCID = cid.reduce((total,[name, value]) => {
          total += value
        }, 0)
      if (totalCID < change) {
        status = 'INSUFFICIENT_FUNDS'
      } else if (totalCID == change) {
        status = 'CLOSED'
      } else {
        status = 'OPEN'
      }

      return {
        status,
        changeDue,
        changed
      }
    }
  }

  total = findAmount('open', change)
  // Here is your change, ma'am.
  return {
    change: total.change,
    status: total.status
  };
}

// Example cash-in-drawer array:
// [["PENNY", 1.01],
// ["NICKEL", 2.05],
// ["DIME", 3.1],
// ["QUARTER", 4.25],
// ["ONE", 90],
// ["FIVE", 55],
// ["TEN", 20],
// ["TWENTY", 60],
// ["ONE HUNDRED", 100]]

checkCashRegister(19.5, 20, [["PENNY", 1.01], ["NICKEL", 2.05], ["DIME", 3.1], ["QUARTER", 4.25], ["ONE", 90], ["FIVE", 55], ["TEN", 20], ["TWENTY", 60], ["ONE HUNDRED", 100]]);